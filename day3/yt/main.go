package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/mbaraa/ytdl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing youtube url")
		os.Exit(1)
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error while parsing the url", err)
		os.Exit(1)
	}

	videoId, err := extractYtId(url)
	if err != nil {
		fmt.Println("Error while extracting the video's ID", err)
		os.Exit(1)
	}

	err = ytdl.DownloadVideo(videoId)
	if err != nil {
		fmt.Println("Error while downloading the video", err)
		os.Exit(1)
	}
}

func extractYtId(videoUrl *url.URL) (string, error) {
	// check for host's validity
	allowedHosts := []string{
		"youtube.com",
		"www.youtube.com",
		"youtu.be",
		"www.youtu.be",
	}
	found := false
	for _, host := range allowedHosts {
		if videoUrl.Hostname() == host {
			found = true
			break
		}
	}
	if !found {
		return "", errors.New("video's url is not supported")
	}

	// extract video url
	if (videoUrl.Hostname() == "youtu.be" ||
		videoUrl.Hostname() == "www.youtu.be") &&
		len(videoUrl.Path) > 1 {
		return videoUrl.Path[1:], nil
	}

	return videoUrl.Query().Get("v"), nil
}
