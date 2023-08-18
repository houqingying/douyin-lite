package favorite_service

import (
	"douyin-lite/repository"
	"fmt"
	"log"
	"sync"
)

type LikeServiceImpl struct {
	VideoService
}

var (
	likeServiceImp      *LikeServiceImpl
	likeServiceInstance sync.Once
)

func NewLikeServImpInstance() *LikeServiceImpl {
	likeServiceInstance.Do(func() {
		likeServiceImp = &LikeServiceImpl{
			VideoService: &VideoServiceImpl{},
		}
	})
	return likeServiceImp
}

// 对接口方法实现
func (*LikeServiceImpl) FavoriteAction(userId int64, videoId int64, actionType int32) error {
	islike, err := repository.IsVideoLikedByUser(userId, videoId)
	log.Print("islike:", islike)
	log.Println("actionType:", actionType)
	// 获取点赞和取消点赞的消息队列
	likeAddMQ := rabbitmq.SimpleLikeAddMQ
	likeDelMQ := rabbitmq.SimpleLikeDelMQ
	if islike == -1 {
		//  更新 redis
		syncLikeRedis(userId, videoId, 1)
		// 消息队列
		err := likeAddMQ.PublishSimple(fmt.Sprintf("%d-%d-%s", userId, videoId, "insert"))
		return err
	}
	//该用户曾对此视频点过赞
	//err = dao.UpdateLikeInfo(userId, videoId, int8(actionType))
	if actionType == 1 {
		syncLikeRedis(userId, videoId, 1)
		err = likeAddMQ.PublishSimple(fmt.Sprintf("%d-%d-%s", userId, videoId, "update"))
	} else {
		syncLikeRedis(userId, videoId, 2)
		err = likeDelMQ.PublishSimple(fmt.Sprintf("%d-%d-%s", userId, videoId, "update"))
	}
	if err != nil {
		log.Print(err.Error() + "Favorite action failed!")
		return err
	} else {
		log.Print("Favorite action succeed!")
	}
	return nil
}
