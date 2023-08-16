package repository

import (
	"fmt"
	"testing"
)

func TestUserDao_CreateUser(t *testing.T) {
	Init()
	userDao.CreateUser("wdp", 0, 0)
	userDao.CreateUser("lwh", 0, 0)
	userDao.CreateUser("cwp", 0, 0)
	userDao.CreateUser("nyh", 0, 0)
	userDao.CreateUser("czx", 0, 0)
	userDao.CreateUser("wjr", 0, 0)
	userDao.CreateUser("cjt", 0, 0)
}

func TestUserDao_CreateRegisterUser(t *testing.T) {
	Init()
	userDao.CreateRegisterUser("wdp", "123")
	userDao.CreateRegisterUser("lwh", "456")
}

func TestUserDao_QueryIsUserExist(t *testing.T) {
	Init()
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
