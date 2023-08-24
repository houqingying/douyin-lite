package entity

import "sync"

type PublishDao struct {
}

var publishDao *PublishDao
var publishOnce sync.Once

func NewPublishDaoInstance() *PublishDao {
	publishOnce.Do(func() {
		publishDao = &PublishDao{}
	})
	return publishDao
}

/*
func (videoService VideoServiceImpl) PublishList(userId int64) ([]models.VideoDVO, error) {
	videoList, err := models.GetVediosByUserId(userId)
	if err != nil {
		return nil, err
	}
	size := len(videoList)
	VideoDVOList := make([]models.VideoDVO, size)
	//创建多个协程并发更新
	var wg sync.WaitGroup
	//接收协程产生的错误
	var err0 error
	for i := range videoList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var userId = videoList[i].AuthorId
			//一定要通过videoService来调用 userSevice
			user, err1 := models.GetUserById(userId)
			if err1 != nil {
				err0 = err1
			}
			var videoDVO models.VideoDVO
			err := copier.Copy(&videoDVO, &videoList[i])
			if err != nil {
				err0 = err1
			}
			videoDVO.Author = user
			VideoDVOList[i] = videoDVO
		}(i)
	}
	wg.Wait()
	//处理协程内的错误
	if err0 != nil {
		return nil, err0
	}
	return VideoDVOList, nil
}
*/
