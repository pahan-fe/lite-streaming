package transcoder

import (
	"context"
	"os"
	"os/exec"
)

type Transcoder struct{}

func (t *Transcoder) TranscodeToHLS(ctx context.Context, inputPath string, outputPath string) error {
	err := os.MkdirAll(outputPath, 0755)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(
		ctx, "ffmpeg", "-i", inputPath, "-c:v", "libx264", "-preset", "fast", "-hls_time", "6", "-hls_segment_filename", outputPath+"/segment_%03d.ts", "-hls_list_size", "0", outputPath+"/index.m3u8")

	return cmd.Run()
}

func NewTranscoder() *Transcoder {
	return &Transcoder{}
}
