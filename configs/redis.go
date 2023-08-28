package configs

import (
	"context"
	"douyin-lite/internal/entity"
	"douyin-lite/pkg/storage"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"sync"
	"time"
)

var RdbUserCount *redis.Client

const RdbUserCountFollowKey = "follow_count:"
const RdbUserCountFollowerKey = "follower_count"

func RedisInit() error {
	RdbUserCount = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	ticker := time.NewTicker(time.Second * 5)
	go StartTimer(ticker)

	_, err := RdbUserCount.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func StartTimer(ticker *time.Ticker) {
	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			fmt.Println("定时任务")
			wg.Add(2)
			go SaveFollowCntToDB(&wg)
			go SaveFollowerCntToDB(&wg)
			wg.Wait()
			fmt.Println("刷新完成")
		}
	}
}

func SaveFollowCntToDB(wg *sync.WaitGroup) error {
	defer wg.Done()
	keys, err := RdbUserCount.Keys(context.Background(), "follow_count:*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		value, err := RdbUserCount.Get(context.Background(), key).Result()
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
		count := entity.Count{
			ID:          id,
			FollowCount: followCnt,
		}

		err = storage.DB.Create(&count).Error
		if err != nil {
			err := storage.DB.Model(&entity.Count{}).Select("follow_count").
				Where("id = ?", id).Update("follow_count", followCnt).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func SaveFollowerCntToDB(wg *sync.WaitGroup) error {
	defer wg.Done()
	keys, err := RdbUserCount.Keys(context.Background(), "follower_count:*").Result()
	if err != nil {
		return err
	}
	for _, key := range keys {
		value, err := RdbUserCount.Get(context.Background(), key).Result()
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
		count := entity.Count{
			ID:            id,
			FollowerCount: followerCnt,
		}

		err = storage.DB.Create(&count).Error
		if err != nil {
			err := storage.DB.Model(&entity.Count{}).Select("follower_count").
				Where("id = ?", id).Update("follower_count", followerCnt).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
