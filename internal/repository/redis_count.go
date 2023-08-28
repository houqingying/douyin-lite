package repository

import (
	"context"
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/storage"
	"fmt"
	"strconv"
	"sync"
)

//type RdbUserCountDao struct {
//}
//
//var rdbUserCountDao *RdbUserCountDao
//var rdbUserCountOnce sync.Once
//
//func NewRdbUserCountDaoInstance() *RdbUserCountDao {
//	rdbUserCountOnce.Do(func() {
//		rdbUserCountDao = &RdbUserCountDao{}
//	})
//	return rdbUserCountDao
//}

func QueryFollowCnt(id int64) (*int64, error) {
	// 查redis
	followCnt, err := queryFollowingCnt(context.Background(), id)
	if err != nil {
		//redis没找到查db
		followCnt, err = entity.NewCountDaoInstance().QueryFollowingCount(id)
		if err != nil {
			return nil, err
		}
		//回写redis
		err := addFollowingCnt(context.Background(), id, *followCnt)
		if err != nil {
			return nil, err
		}
	}
	return followCnt, nil
}

func QueryFollowerCnt(id int64) (*int64, error) {
	// 查redis
	followerCnt, err := queryFollowerCnt(context.Background(), id)
	if err != nil {
		//redis没找到查db
		followerCnt, err = entity.NewCountDaoInstance().QueryFollowerCount(id)
		if err != nil {
			return nil, err
		}
		//回写redis
		err := addFollowerCnt(context.Background(), id, *followerCnt)
		if err != nil {
			return nil, err
		}
	}
	return followerCnt, nil
}

func IncFollowingCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follow_count:%d", hostId)
	_, err := storage.RdbUserCount.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func DecFollowingCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follow_count:%d", hostId)
	_, err := storage.RdbUserCount.Decr(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func IncFollowerCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	_, err := storage.RdbUserCount.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func DecFollowerCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	_, err := storage.RdbUserCount.Decr(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func SaveFollowCntToDB(wg *sync.WaitGroup) error {
	defer wg.Done()
	keys, err := storage.RdbUserCount.Keys(context.Background(), "follow_count:*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		value, err := storage.RdbUserCount.Get(context.Background(), key).Result()
		if err != nil {
			return err
		}
		followCnt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		idStr := key[len("follow_count:"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}
		err = entity.NewCountDaoInstance().SaveFollowingCount(id, followCnt)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveFollowerCntToDB(wg *sync.WaitGroup) error {
	defer wg.Done()
	keys, err := storage.RdbUserCount.Keys(context.Background(), "follower_count:*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		value, err := storage.RdbUserCount.Get(context.Background(), key).Result()
		if err != nil {
			return err
		}
		followerCnt, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		idStr := key[len("follower_count:"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}
		err = entity.NewCountDaoInstance().SaveFollowerCount(id, followerCnt)
		if err != nil {
			return err
		}
	}
	return nil
}

// private
func queryFollowerCnt(ctx context.Context, hostId int64) (*int64, error) {
	key := fmt.Sprintf("follower_count:%d", hostId)
	res, err := storage.RdbUserCount.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	atoi, err := strconv.Atoi(res)
	fc := int64(atoi)
	if err != nil {
		return nil, err
	}
	fmt.Printf("返回的值为 %d\n", atoi)
	return &fc, nil
}

func queryFollowingCnt(ctx context.Context, hostId int64) (*int64, error) {
	key := fmt.Sprintf("follow_count:%d", hostId)
	res, err := storage.RdbUserCount.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	atoi, err := strconv.Atoi(res)
	fc := int64(atoi)
	if err != nil {
		return nil, err
	}
	fmt.Printf("返回的值为 %d\n", atoi)
	return &fc, nil
}

func addFollowingCnt(ctx context.Context, hostId int64, cnt int64) error {
	key := fmt.Sprintf("follow_count:%d", hostId)
	_, err := storage.RdbUserCount.Set(ctx, key, cnt, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func addFollowerCnt(ctx context.Context, hostId int64, cnt int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	_, err := storage.RdbUserCount.Set(ctx, key, cnt, 0).Result()
	if err != nil {
		return err
	}
	return nil
}
