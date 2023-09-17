package relation_service

import (
	"errors"

	"gorm.io/gorm"

	"douyin-lite/internal/entity"
	"douyin-lite/internal/service/message_service"
	"douyin-lite/internal/service/user_service"
)

type FriendListInfo struct {
	FrinedUserInfoList []*FriendUserInfo `json:"user_list,omitempty"`
}

type FriendUserInfo struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowingCount  int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
	Message         string `json:"message,omitempty"` // 和该好友的最新聊天消息
	MsgType         int64  `json:"msgType"`           // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func QueryFriendListInfo(hostId int64) (*FriendListInfo, error) {
	return NewQueryFriendListInfoFlow(hostId).Do()
}

type QueryFriendListInfoFlow struct {
	hostId         int64
	friendListInfo *FriendListInfo

	userinfoList []*FriendUserInfo
}

func NewQueryFriendListInfoFlow(hostId int64) *QueryFriendListInfoFlow {
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
	friendIdList, err := FollowingDao.QueryFriendIdListById(f.hostId)
	if err != nil {
		return errors.New("DB Find FriendIdList Error")
	}
	friendUserInfoList, err := user_service.QueryUserInfoList(f.hostId, &friendIdList)
	if err != nil {
		return err
	}

	var friendInfoList = make([]*FriendUserInfo, len(friendUserInfoList))
	for i, friend := range friendUserInfoList {

		var message string
		var msgType int64
		msgInfo, err := message_service.QueryLastMessage(f.hostId, friend.ID)

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("DB Find Message Error")
			}
			// 如果查找最新消息时出现错误，且为RecordNotFound，设置为""空串, msgType设为1
			message = ""
			msgType = 1
		} else {
			// 根据发送方向设置msgType
			message = msgInfo.Content
			if msgInfo.FromUserID == f.hostId {
				msgType = 1
			} else {
				msgType = 0
			}
		}

		friendInfoList[i] = &FriendUserInfo{
			ID:              friend.ID,
			Name:            friend.Name,
			Avatar:          friend.Avatar,
			BackgroundImage: friend.BackgroundImage,
			Signature:       friend.Signature,
			FollowingCount:  friend.FollowingCount,
			FollowerCount:   friend.FollowerCount,
			IsFollow:        friend.IsFollow,
			TotalFavorited:  friend.TotalFavorited,
			WorkCount:       friend.WorkCount,
			FavoriteCount:   friend.FavoriteCount,
			Message:         message,
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
