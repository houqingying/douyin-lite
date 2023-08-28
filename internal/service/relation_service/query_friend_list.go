package relation_service

import (
	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/message_service"
	"errors"
)

type FriendListInfo struct {
	FrinedUserInfoList []*FriendUserInfo `json:"user_list,omitempty"`
}

type FriendUserInfo struct {
	ID              int64  `json:"id,omitempty"`               // 用户id
	Name            string `json:"name,omitempty"`             // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar,omitempty"`           //用户头像
	BackgroundImage string `json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited,omitempty"`  //获赞数量
	WorkCount       int64  `json:"work_count,omitempty"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count,omitempty"`   //点赞数量
	Message         string `json:"message,omitempty"`          // 和该好友的最新聊天消息
	MsgType         int64  `json:"msgType,omitempty"`          // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func QueryFriendListInfo(hostId uint) (*FriendListInfo, error) {
	return NewQueryFriendListInfoFlow(hostId).Do()
}

type QueryFriendListInfoFlow struct {
	hostId         uint
	friendListInfo *FriendListInfo

	userinfoList []*FriendUserInfo
}

func NewQueryFriendListInfoFlow(hostId uint) *QueryFriendListInfoFlow {
	return &QueryFriendListInfoFlow{
		hostId: hostId,
	}
}

var FollowingDao = entity.NewFollowingDaoInstance()

func (f *QueryFriendListInfoFlow) Do() (*FriendListInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareFriendInfo()
	if err != nil {
		return nil, err
	}
	err = f.packFriendInfo()
	if err != nil {
		return nil, err
	}
	return f.friendListInfo, nil
}

func (f *QueryFriendListInfoFlow) checkParam() error {
	if f.hostId <= 0 {
		return errors.New("host id should be larger than 0")
	}
	return nil
}

func (f *QueryFriendListInfoFlow) prepareFriendInfo() error {
	friendList, err := FollowingDao.QueryFriendById(f.hostId)
	if err != nil {
		return errors.New("DB Find Friend Error")
	}

	var friendInfoList = make([]*FriendUserInfo, len(friendList))
	for i, friend := range friendList {
		var msgType int64
		msgInfo, err := message_service.QueryLastMessage(f.hostId, uint(friend.ID))
		if err != nil {
			return errors.New("DB Find Message Error")
		}
		if uint(msgInfo.FromUserID) == f.hostId {
			msgType = 1
		} else {
			msgType = 0
		}
		friendInfoList[i] = &FriendUserInfo{
			ID:              friend.ID,
			Name:            friend.Name,
			Avatar:          friend.Avatar,
			BackgroundImage: friend.BackgroundImage,
			Signature:       friend.Signature,
			FollowCount:     friend.FollowingCount,
			FollowerCount:   friend.FollowerCount,
			IsFollow:        true,
			TotalFavorited:  friend.TotalFavorited,
			WorkCount:       friend.WorkCount,
			FavoriteCount:   friend.FavoriteCount,
			Message:         msgInfo.Content,
			MsgType:         msgType,
		}
	}

	f.userinfoList = friendInfoList

	return nil
}

func (f *QueryFriendListInfoFlow) packFriendInfo() error {
	f.friendListInfo = &FriendListInfo{
		FrinedUserInfoList: f.userinfoList,
	}
	return nil
}
