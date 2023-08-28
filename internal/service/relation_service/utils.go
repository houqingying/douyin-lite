package relation_service

import (
	"context"
	"douyin-lite/internal/entity"
	"douyin-lite/middleware"
)

func GetFollowAndFollowerCnt(id int64) (*int64, *int64, error) {
	// 查FollowingCnt和FollowerCnt
	followCnt, err := middleware.QueryFollowingCnt(context.Background(), id)
	if err != nil {
		followCnt, err = entity.NewCountDaoInstance().QueryFollowingCount(id)
		if err != nil {
			return nil, nil, err
		}
		//回写redis
		middleware.AddFollowingCnt(context.Background(), id, *followCnt)
	}
	followerCnt, err := middleware.QueryFollowerCnt(context.Background(), id)
	if err != nil {
		followerCnt, err = entity.NewCountDaoInstance().QueryFollowerCount(id)
		if err != nil {
			return nil, nil, err
		}
		//回写redis
		middleware.AddFollowerCnt(context.Background(), id, *followerCnt)
	}
	return followCnt, followerCnt, nil
}
