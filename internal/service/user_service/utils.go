package user_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/repository"
)

// QueryUserInfoList @title
// @description   该方法传入一组id(idList), 查询一组id对应的userInfo (userInfo的结构体定义在user_service/query_user_info)
// @param     hostId      int64         需要查询的hostId, isFollow表示hostId是否对某一个id进行了关注(isFollow)
// @param     idList    *[]int64        需要查询的一组idList
// @return    			 []*UserInfo    返回查询的idList对应的一组UserInfo的地址
func QueryUserInfoList(hostId int64, idList *[]int64) ([]*entity.UserInfo, error) {
	var userInfoList = make([]*entity.UserInfo, len(*idList))
	for i, id := range *idList {
		userInfo, err := QueryAUserInfo1(hostId, id)
		if err != nil {
			return nil, err
		}
		userInfoList[i] = userInfo
	}
	return userInfoList, nil
}

// QueryAUserInfo1 @title
// @description   该方法传入一个id, 查询该id对应的userInfo
// @param     hostId      int64         需要查询的hostId, isFollow表示hostId是否对某一个id进行了关注(isFollow)
// @param     id          int64         需要查询的id
// @return    			  *UserInfo     返回查询的id对应的UserInfo地址
func QueryAUserInfo1(hostId int64, id int64) (*entity.UserInfo, error) {
	user, err := entity.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	followCnt, err := repository.QueryFollowCnt(id)
	if err != nil {
		return nil, err
	}
	followerCnt, err := repository.QueryFollowerCnt(id)
	if err != nil {
		return nil, err
	}
	isFollow, err := entity.NewFollowingDaoInstance().QueryisFollow(hostId, id)
	if err != nil {
		return nil, err
	}
	userInfo := &entity.UserInfo{
		ID:              user.ID,
		Name:            user.Name,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		FollowingCount:  *followCnt,
		FollowerCount:   *followerCnt,
		IsFollow:        isFollow,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return userInfo, nil
}

// QueryAUserInfo2 @title
// @description   该方法传入一个id, 查询该id对应的userInfo. 该方法没有参数hostId, isFollow默认为true
// @param     id          int64         需要查询的id
// @return    			  *UserInfo     返回查询的id对应的UserInfo地址
func QueryAUserInfo2(id int64) (*entity.UserInfo, error) {
	user, err := entity.NewUserDaoInstance().QueryUserById(id)
	if err != nil {
		return nil, err
	}
	followCnt, err := repository.QueryFollowCnt(id)
	if err != nil {
		return nil, err
	}
	followerCnt, err := repository.QueryFollowerCnt(id)
	if err != nil {
		return nil, err
	}
	isFollow := true
	if err != nil {
		return nil, err
	}
	userInfo := &entity.UserInfo{
		ID:              user.ID,
		Name:            user.Name,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		FollowingCount:  *followCnt,
		FollowerCount:   *followerCnt,
		IsFollow:        isFollow,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
	return userInfo, nil
}
