package repository

import (
	"context"
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/storage"
	"errors"
	"strconv"
)

func QueryIsFollow(hostId int64, guestId int64) (bool, error) {
	isFollow, err := queryIsFollow(hostId, guestId)
	if err != nil {
		isFollow, err = entity.NewFollowingDaoInstance().QueryisFollow(hostId, guestId)
		if err != nil {
			return false, err
		}
		// 回写redis
		err = PostFollow(hostId, guestId)
		if err != nil {
			return false, err
		}
	}
	return isFollow, nil
}

func PostFollow(hostId int64, guestId int64) error {
	isFollow, err := queryIsFollow(hostId, guestId)
	if err != nil {
		return err
	}
	isFollower, err := queryIsFollower(guestId, hostId)
	if err != nil {
		return err
	}
	if isFollow || isFollower {
		return errors.New("已经是关注关系或者关注关系产生了错误")
	}
	hostIdStr := strconv.FormatInt(hostId, 10)
	guestIdStr := strconv.FormatInt(guestId, 10)
	err = storage.RdbFollow.SAdd(context.Background(), hostIdStr, guestId).Err()
	if err != nil {
		return err
	}
	err = storage.RdbFollower.SAdd(context.Background(), guestIdStr, hostId).Err()
	if err != nil {
		return err
	}
	return nil
}

func PostUnfollow(hostId int64, guestId int64) error {
	isFollow, err := queryIsFollow(hostId, guestId)
	if err != nil {
		return err
	}
	isFollower, err := queryIsFollower(guestId, hostId)
	if err != nil {
		return err
	}
	if !isFollow || !isFollower {
		return errors.New("并不是关注关系或者关注关系产生了错误")
	}
	hostIdStr := strconv.FormatInt(hostId, 10)
	guestIdStr := strconv.FormatInt(guestId, 10)
	err = storage.RdbFollow.SRem(context.Background(), hostIdStr, guestId).Err()
	if err != nil {
		return err
	}
	err = storage.RdbFollower.SRem(context.Background(), guestIdStr, hostId).Err()
	if err != nil {
		return err
	}
	return nil
}

// func SaveFollowRelationToDB() error {
//
// }
func queryIsFollow(hostId int64, guestId int64) (bool, error) {
	hostIdStr := strconv.FormatInt(hostId, 10)
	result, err := storage.RdbFollow.SIsMember(context.Background(), hostIdStr, guestId).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

func queryIsFollower(hostId int64, guestId int64) (bool, error) {
	hostIdStr := strconv.FormatInt(hostId, 10)
	result, err := storage.RdbFollower.SIsMember(context.Background(), hostIdStr, guestId).Result()
	if err != nil {
		return false, err
	}
	return result, err
}
