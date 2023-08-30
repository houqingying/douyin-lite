package configs

import (
	"context"
	"douyin-lite/internal/repository"
	"douyin-lite/pkg/storage"
	"sync"
	"time"
)

const RdbUserCountFollowKey = "follow_count:"
const RdbUserCountFollowerKey = "follower_count:"
const TimeInterval = time.Second * 5

func RedisInit() error {
	err := storage.InitRedis()
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
	err = storage.RdbUserCount.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	err = storage.RdbFollower.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
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
	return nil
}
