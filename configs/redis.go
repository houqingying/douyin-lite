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
	go StartTimer(ticker)
	_, err = storage.RdbUserCount.Ping(context.Background()).Result()
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
			go repository.SaveFollowCntToDB(&wg)
			go repository.SaveFollowerCntToDB(&wg)
			wg.Wait()
			fmt.Println("刷新完成")
		}
	}
}
