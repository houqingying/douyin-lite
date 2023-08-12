package main

import (
	"github.com/houqingying/douyin-lite/pkg/config"
	"k8s.io/klog"
)

func main() {
	if err := config.Setup(); err != nil {
		klog.Fatalf("config.Setup() error: %s", err)
	}
}
