package ffmpeg

import (
	"bytes"
	"fmt"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// ReadFrameAsJpeg 从视频文件中读取指定帧并将其保存为 JPEG 图像。
// inFileName 是输入视频文件的路径。
// outImagePath 是保存 JPEG 图像的输出路径。
// frameNum 是要提取的帧的帧号。
// 返回可能的错误。
func ReadFrameAsJpeg(inFileName, outImagePath string, frameNum int) (err error) {
	// 创建一个用于存储 ffmpeg 输出的内存缓冲区。
	reader := bytes.NewBuffer(nil)

	// 使用 ffmpeg 库执行以下操作：
	// 1. 输入视频文件 inFileName。
	// 2. 使用 select 过滤器选择帧号大于等于 frameNum 的帧。
	// 3. 输出一帧图像到标准输出，格式为 JPEG。
	if err = ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run(); err != nil {
		return err
	}

	// 使用 imaging 库解码从 ffmpeg 输出中读取的 JPEG 数据。
	img, err := imaging.Decode(reader)
	if err != nil {
		return err
	}

	// 将解码后的图像保存到指定的输出路径。
	err = imaging.Save(img, outImagePath)
	if err != nil {
		return err
	}

	// 返回 nil 表示没有错误。
	return nil
}
