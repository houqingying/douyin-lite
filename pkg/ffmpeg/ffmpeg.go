package ffmpeg

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

// ReadFrameAsJpeg 从视频文件中读取指定帧并将其保存为 JPEG 图像。
// inFileName 是输入视频文件的路径。
// outImagePath 是保存 JPEG 图像的输出路径。
// frameNum 是要提取的帧的帧号。
// 返回可能的错误。
func ReadFrameAsJpeg(inFileName, outImagePath string) (err error) {
	// SSH连接配置
	serverAddr := "47.102.185.103:22"
	username := "root"
	password := "a#A6@FLVaDnsyH"

	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不验证服务器主机密钥
	}

	// 连接SSH服务器
	client, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		log.Fatalf("无法连接到SSH服务器：%v", err)
	}
	defer client.Close()

	// 执行FFmpeg截图命令
	ffmpegCommand := fmt.Sprintf("ffmpeg -i %s -ss 00:00:01 -vframes 1 %s", inFileName, outImagePath)

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("无法创建SSH会话：%v", err)
		return err
	}
	defer session.Close()

	_, err = session.CombinedOutput(ffmpegCommand)
	if err != nil {
		log.Fatalf("FFmpeg截图命令执行失败：%v", err)
		return err
	}

	// 返回 nil 表示没有错误。
	return nil
}
