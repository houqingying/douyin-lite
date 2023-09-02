package entity

import "sync"

type PublishDao struct {
}

var publishDao *PublishDao
var publishOnce sync.Once

func NewPublishDaoInstance() *PublishDao {
	publishOnce.Do(func() {
		publishDao = &PublishDao{}
	})
	return publishDao
}
