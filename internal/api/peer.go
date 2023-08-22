package api

import (
	"douyin-lite/pkg/fastdfs"
	"github.com/gin-gonic/gin"
)

// getPeersUrl 获取PeersUrl
func getPeersUrl(ctx *gin.Context) (string, error) {
	peers := fastdfs.NewPeers()

	if peers.GroupName != "" {
		return peers.ServerAddress + "/" + peers.GroupName, nil
	}
	return peers.ServerAddress, nil
}

// getPeers 获取Peers
func getPeers(ctx *gin.Context) (*fastdfs.Peers, error) {
	return fastdfs.NewPeers(), nil
}

// getUser 获取当前用户
func getUser(ctx *gin.Context) (*fastdfs.PeerUser, error) {
	return fastdfs.NewPeerUser(), nil
}

// getShowUrl 获取ShowUrl
func getShowUrl(ctx *gin.Context) string {
	peers, _ := getPeers(ctx)
	showUrl := ""
	if peers.ShowAddress == "" {
		if peers.GroupName == "" {
			showUrl = peers.ServerAddress
		} else {
			showUrl = peers.ServerAddress + "/" + peers.GroupName
		}
	} else {
		if peers.GroupName == "" {
			showUrl = peers.ShowAddress
		} else {
			showUrl = peers.ShowAddress + "/" + peers.GroupName
		}
	}
	return showUrl
}

// getShowUrlNotGroup 获取ShowUrl
func getShowUrlNotGroup(ctx *gin.Context) string {
	peers, _ := getPeers(ctx)
	showUrl := ""
	if peers.ShowAddress == "" {
		showUrl = peers.ServerAddress
	} else {
		showUrl = peers.ShowAddress
	}
	return showUrl
}
