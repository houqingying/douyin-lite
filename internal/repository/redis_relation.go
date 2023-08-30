package repository

import (
	"context"
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/storage"
	"errors"
	"strconv"
	"strings"
	"sync"
)

// RelationScanNum 每次定时任务Scan从redis删除写入到mysql的数量
const RelationScanNum = 5
const LargeSetNum = 100
const DeleteOneSetNum = 10

func QueryIsFollow(hostId int64, guestId int64) (bool, error) {
	isFollow, err := queryIsFollow(hostId, guestId)
	if err != nil {
		isFollow, err = entity.NewFollowingDaoInstance().QueryisFollow(hostId, guestId)
		if err != nil {
			return false, err
		}
		// 如果存在关注关系,回写redis
		if isFollow {
			err = FollowAction(hostId, guestId)
			if err != nil {
				return false, err
			}
		}
	}
	return isFollow, nil
}

func FollowAction(hostId int64, guestId int64) error {
	hostIdStr := strconv.FormatInt(hostId, 10)
	guestIdStr := strconv.FormatInt(guestId, 10)
	followStateIdStr := hostIdStr + ":1"
	unfollowStateIdStr := hostIdStr + ":0"
	isUnfollowRes, err := storage.RdbFollower.SIsMember(context.Background(), guestIdStr, unfollowStateIdStr).Result()
	if isUnfollowRes {
		//先删
		err := storage.RdbFollower.SRem(context.Background(), guestIdStr, unfollowStateIdStr).Err()
		if err != nil {
			return err
		}
	}
	err = storage.RdbFollower.SAdd(context.Background(), guestIdStr, followStateIdStr).Err()
	if err != nil {
		return err
	}
	return nil
}

func UnfollowAction(hostId int64, guestId int64) error {
	hostIdStr := strconv.FormatInt(hostId, 10)
	guestIdStr := strconv.FormatInt(guestId, 10)
	unfollowStateIdStr := hostIdStr + ":0"
	followStateIdStr := hostIdStr + ":1"
	isFollowRes, err := storage.RdbFollower.SIsMember(context.Background(), guestIdStr, followStateIdStr).Result()
	if err != nil {
		return err
	}
	if isFollowRes {
		// 先删
		err := storage.RdbFollower.SRem(context.Background(), guestIdStr, followStateIdStr).Err()
		if err != nil {
			return err
		}
	}
	// 增加
	err = storage.RdbFollower.SAdd(context.Background(), guestIdStr, unfollowStateIdStr).Err()
	if err != nil {
		return err
	}
	return nil
}

func SaveFollowRelationToDB(wg *sync.WaitGroup, cursor *uint64) error {
	defer wg.Done()
	// 得到redis一部分键, scan防止阻塞
	keys, res, err := storage.RdbFollower.Scan(context.Background(), *cursor, "*", RelationScanNum).Result()
	if err != nil {
		return err
	}
	*cursor = res
	for _, key := range keys {
		guestId, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return err
		}
		setNum, err := storage.RdbFollower.SCard(context.Background(), key).Result()
		if err != nil {
			return err
		}
		// 如果一个用户的粉丝非常庞大, 大概率是热点用户, 暂时不清除
		if setNum > LargeSetNum {
			continue
		}
		// 从集合总拿出一定数量的元素
		stateList, err := storage.RdbFollower.SPopN(context.Background(), key, DeleteOneSetNum).Result()
		if err != nil {
			return err
		}
		for _, state := range stateList {
			hostId, flag, err := split(state)
			if err != nil {
				return err
			}
			if flag == 0 {
				//数据库删除关注关系
				entity.NewFollowingDaoInstance().DeleteFollowing(hostId, guestId)
			} else {
				//数据库增加关注关系
				entity.NewFollowingDaoInstance().CreateFollowing(hostId, guestId)
			}
		}
	}
	return nil
}

func split(state string) (int64, int64, error) {
	parts := strings.Split(state, ":")
	if len(parts) != 2 {
		return -1, -1, errors.New("split错误")
	}
	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return -1, -1, err
	}
	flag, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return -1, -1, err
	}
	return id, flag, nil
}
func queryIsFollow(hostId int64, guestId int64) (bool, error) {
	hostIdStr := strconv.FormatInt(hostId, 10)
	guestIdStr := strconv.FormatInt(guestId, 10)
	followStateIdStr := hostIdStr + ":1"
	// 查guest的set里面有没有host
	isFollowRes, err := storage.RdbFollower.SIsMember(context.Background(), guestIdStr, followStateIdStr).Result()
	if err != nil {
		return false, err
	}
	if isFollowRes {
		return true, nil
	}
	unfollowStateIdStr := hostIdStr + ":0"
	isUnfollowRes, err := storage.RdbFollower.SIsMember(context.Background(), guestIdStr, unfollowStateIdStr).Result()
	if err != nil {
		return false, err
	}
	if isUnfollowRes {
		return false, nil
	}
	return false, errors.New("redis Not Found")
}
