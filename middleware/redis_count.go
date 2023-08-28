package middleware

import (
	"context"
	"douyin-lite/configs"
	"fmt"
	"strconv"
)

func IncFollowingCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("following_cnt:%d", hostId)
	res, err := configs.RdbUserCount.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	fmt.Printf("增加后的值为%d\n", res)
	return nil
}

func DecFollowingCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("following_cnt:%d", hostId)
	res, err := configs.RdbUserCount.Decr(ctx, key).Result()
	if err != nil {
		return err
	}
	fmt.Printf("增加后的值为%d\n", res)
	return nil
}

func IncFollowerCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	res, err := configs.RdbUserCount.Incr(ctx, key).Result()
	if err != nil {
		return err
	}
	fmt.Printf("增加后的值为%d\n", res)
	return nil
}

func DecFollowerCnt(ctx context.Context, hostId int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	res, err := configs.RdbUserCount.Decr(ctx, key).Result()
	if err != nil {
		return err
	}
	fmt.Printf("增加后的值为%d\n", res)
	return nil
}

func QueryFollowerCnt(ctx context.Context, hostId int64) (*int64, error) {
	key := fmt.Sprintf("follower_count:%d", hostId)
	res, err := configs.RdbUserCount.Get(ctx, key).Result()
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

func QueryFollowingCnt(ctx context.Context, hostId int64) (*int64, error) {
	key := fmt.Sprintf("follow_count:%d", hostId)
	res, err := configs.RdbUserCount.Get(ctx, key).Result()
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

func AddFollowingCnt(ctx context.Context, hostId int64, cnt int64) error {
	key := fmt.Sprintf("follow_count:%d", hostId)
	_, err := configs.RdbUserCount.Set(ctx, key, cnt, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func AddFollowerCnt(ctx context.Context, hostId int64, cnt int64) error {
	key := fmt.Sprintf("follower_count:%d", hostId)
	_, err := configs.RdbUserCount.Set(ctx, key, cnt, 0).Result()
	if err != nil {
		return err
	}
	return nil
}
