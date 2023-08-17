package test

import (
	"douyin-lite/repository"
	"fmt"
	"testing"
)

func TestUserDao_CreateUser(t *testing.T) {
	repository.Init()
	userDao := repository.NewUserDaoInstance()
	userDao.CreateUser("wdp", 0, 0)
	userDao.CreateUser("lwh", 0, 0)
	userDao.CreateUser("cwp", 0, 0)
	userDao.CreateUser("nyh", 0, 0)
	userDao.CreateUser("czx", 0, 0)
	userDao.CreateUser("wjr", 0, 0)
	userDao.CreateUser("cjt", 0, 0)
}

func TestUserDao_CreateRegisterUser(t *testing.T) {
	repository.Init()
	userDao := repository.NewUserDaoInstance()
	userDao.CreateRegisterUser("wdp", "123")
	userDao.CreateRegisterUser("lwh", "456")
}

func TestUserDao_QueryIsUserExist(t *testing.T) {
	repository.Init()
	userDao := repository.NewUserDaoInstance()
	isExist, err := userDao.QueryIsUserExist("wangdongdong")
	if err != nil {
		if isExist == false {
			fmt.Println("Not Exist")
			return
		} else {
			panic("存在异常")
		}
	}
	fmt.Println("Exist")
}

func TestUserDao_QueryIsUserLogin(t *testing.T) {
	repository.Init()
	userDao := repository.NewUserDaoInstance()
	qUser, err := userDao.QueryLoginUser("mao122", "123")
	if err != nil {
		panic(err)
	}
	fmt.Println(qUser)
}

func TestUserDao_QueryUserById(t *testing.T) {
	repository.Init()
	userDao := repository.NewUserDaoInstance()
	qUser, err := userDao.QueryUserById(21)
	if err != nil {
		panic(err)
	}
	fmt.Println(qUser)
}
