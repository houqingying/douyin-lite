package snowflake

import (
	"fmt"
	"testing"
)

func TestSnakeFlow(t *testing.T) {
	// 初始化雪花算法
	err := InitSnowflakeNode("2023-07-01", 1)
	if err != nil {
		panic(err)
	}
	ID1 := GenerateID()
	fmt.Println(ID1)
	ID2 := GenerateID()
	fmt.Println(ID2)
	ID3 := GenerateID()
	fmt.Println(ID3)
}
