package repository

import "testing"

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
