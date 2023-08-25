package ffmpeg

import (
	"bytes"
	"fmt"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ReadFrameAsJpeg(inFileName, outImagePath string, frameNum int) (err error) {
	reader := bytes.NewBuffer(nil)
	if err = ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run(); err != nil {
		return err
	}

	img, err := imaging.Decode(reader)
	if err != nil {
		return err
	}
	err = imaging.Save(img, outImagePath)
	if err != nil {
		return err
	}
	return nil
}
