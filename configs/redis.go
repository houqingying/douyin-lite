package configs

import (
	"context"
	"douyin-lite/internal/repository"
	"douyin-lite/pkg/storage"
	"fmt"
	"sync"
	"time"
)

const RdbUserCountFollowKey = "follow_count:"
const RdbUserCountFollowerKey = "follower_count:"

func RedisInit() error {
	err := storage.InitRedis()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time.Second * 5)
	go func() {
		err := StartTimer(ticker)
		if err != nil {
			panic("定时任务启动失败")
		}
	}()
	_, err = storage.RdbUserCount.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func StartTimer(ticker *time.Ticker) error {
	var wg sync.WaitGroup
	for {
		select {
		case <-ticker.C:
			var err error
			fmt.Println("定时任务")
			wg.Add(2)
			go func() {
				err = repository.SaveFollowCntToDB(&wg)
			}()
			go func() {
				err = repository.SaveFollowerCntToDB(&wg)
			}()
			wg.Wait()
			if err != nil {
				return err
			}
			fmt.Println("刷新完成")
			return nil
		}
	}
}
