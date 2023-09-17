package configs

import (
	"context"
	conf "douyin-lite/configs/locales"
	"douyin-lite/internal/repository"
	"douyin-lite/pkg/storage"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const RdbUserCountFollowKey = "follow_count:"
const RdbUserCountFollowerKey = "follower_count:"
const TimeInterval = time.Second * 5

func RedisInit() error {
	mConfigDB0 := conf.Config.Redis["db0"]
	storage.RdbUserCount = redis.NewClient(&redis.Options{
		Addr:     mConfigDB0.RedisAddr,
		Password: mConfigDB0.RedisPassword,
		DB:       mConfigDB0.RedisDbName,
	})
	mConfigDB1 := conf.Config.Redis["db1"]
	storage.RdbFollower = redis.NewClient(&redis.Options{
		Addr:     mConfigDB1.RedisAddr,
		Password: mConfigDB1.RedisPassword,
		DB:       mConfigDB1.RedisDbName,
	})
	err := storage.RdbUserCount.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	err = storage.RdbFollower.Ping(context.Background()).Err()
	if err != nil {
		return err
	}

	// 每隔TimeInterval触发一次定时任务
	ticker := time.NewTicker(TimeInterval)
	go func() {
		err := StartTimer(ticker)
		if err != nil {
			panic("定时任务启动失败")
		}
	}()
	return nil
}

func StartTimer(ticker *time.Ticker) error {
	var wg sync.WaitGroup
	var cursor1 uint64
	var cursor2 uint64
	var cursor3 uint64
	for {
		select {
		case <-ticker.C:
			var err error
			//定时任务开启
			wg.Add(3)
			// 定时任务: 每隔一段时间将Redis的FollowCnt写到DB
			go func() {
				err = repository.SaveFollowCntToDB(&wg, &cursor1)
				if err != nil {
					panic(err)
				}
			}()
			// 定时任务: 每隔一段时间将Redis的FollowerCnt写到DB
			go func() {
				err = repository.SaveFollowerCntToDB(&wg, &cursor2)
				if err != nil {
					panic(err)
				}
			}()
			// 定时任务: 每隔一段时间将Redis的FollowRelation写到DB
			go func() {
				err = repository.SaveFollowRelationToDB(&wg, &cursor3)
				if err != nil {
					panic(err)
				}
			}()
			wg.Wait()
			//定时任务结束
		}
	}
}
