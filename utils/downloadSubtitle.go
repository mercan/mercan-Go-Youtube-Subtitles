package utils

import (
	"log"
	"os/exec"
)

const (
	noCheckCertificate = "--no-check-certificate"
	skipDownload       = "--skip-download"
	writeAutoSub       = "--write-auto-sub"
	subLanguage        = "--sub-lang=tr"
	outputFlag         = "-o"
	outputPath         = "subtitles/%(id)s.%(ext)s"
	subFormatExt       = "--sub-format=ttml"
)

func DownloadSubtitle(videoID string) error {
	cmd := exec.Command("youtube-dl",
		noCheckCertificate,
		skipDownload,
		writeAutoSub,
		subLanguage,
		subFormatExt,
		outputFlag,
		outputPath,
		"--",
		videoID,
	)

	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	log.Printf("Subtitle Downloaded Successfully VideoID: %s\n", videoID)
	return nil
}
