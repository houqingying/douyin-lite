package configs

import (
	conf "douyin-lite/configs/locales"
	"douyin-lite/pkg/ffmpeg"
)

// 错误消息常量
const (
	ErrReadFrameFailure = "读取视频帧失败"
)

func FfmpegInit() {
	ffmpeg.NewFfmpegClient(
		conf.Config.Ffmpeg.ServerAddr,
		conf.Config.Ffmpeg.Username,
		conf.Config.Ffmpeg.Password,
	)
}
