package follow_service

import (
	"douyin-lite/repository"
	"errors"
	"fmt"
)

const (
	FOLLOW = iota + 1
	CANCEL
)

var (
	ErrIvdAct    = errors.New("undefined action")
	ErrIvdFolUsr = errors.New("user not exist")
)

func FollowAction(HostId, GuestId, actionType uint) error {
	return NewPostFollowActionFlow(HostId, GuestId, actionType).Do()
}

type PostFollowActionFlow struct {
	userId     uint
	userToId   uint
	actionType uint
}

func NewPostFollowActionFlow(userId, userToId, actionType uint) *PostFollowActionFlow {
	return &PostFollowActionFlow{userId: userId, userToId: userToId, actionType: actionType}
}

func (p *PostFollowActionFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}
	if err = p.action(); err != nil {
		return err
	}
	return nil
}

func (p *PostFollowActionFlow) checkNum() error {
	//check userToId
	isExist, err := repository.NewUserDaoInstance().QueryIsUserExistById(p.userToId)
	if err != nil {
		return err
	}
	if !isExist {
		return ErrIvdFolUsr
	}
	if p.actionType != FOLLOW && p.actionType != CANCEL {
		return ErrIvdAct
	}
	// can't follow self
	if p.userId == p.userToId {
		return ErrIvdAct
	}
	return nil
}

func (p *PostFollowActionFlow) action() error {
	exist, err := repository.NewFollowingDaoInstance().QueryisFollow(p.userId, p.userToId)
	if err != nil {
		return err
	}

	switch p.actionType {
	case FOLLOW:
		if exist {
			return fmt.Errorf("relation already exist")
		}
		err = repository.NewFollowingDaoInstance().FollowAction(p.userId, p.userToId)
		if err != nil {
			return err
		}
	case CANCEL:
		if !exist {
			return fmt.Errorf("relation not exist")
		}
		err = repository.NewFollowingDaoInstance().UnfollowAction(p.userId, p.userToId)
		if err != nil {
			return err
		}
	}

	return nil
}
