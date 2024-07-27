package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"ytcli/cfmt"

	"github.com/mbaraa/ytdl"
	"github.com/mbaraa/ytscrape"
)

func main() {
	if len(os.Args) < 2 {
		cfmt.Red().Bold().Println("Enter a search term!")
		os.Exit(1)
	}

	searchQuery := os.Args[1]
	results, err := ytscrape.Search(searchQuery)
	if err != nil {
		cfmt.Red().Bold().Println("Error occurred when searching youtube ", err)
		os.Exit(1)
	}

	for i := len(results) - 1; i >= 0; i-- {
		fmt.Printf("%s %s\n", cfmt.Blue().Sprint(i+1), resultItem(results[i]))
	}

	cfmt.Green().Bold().Print("==> ")
	fmt.Println("Video to download (eg: 1, 2, 3...)")
	cfmt.Green().Bold().Print("==> ")
	var choice int
	fmt.Scan(&choice)

	if choice-1 > len(results) || choice < 0 {
		cfmt.Red().Println("INVALID CHOICE!!!!!!")
		os.Exit(1)
	}

	chosenResult := results[choice-1]

	fmt.Println("Chose what to do:")
	cfmt.Cyan().Println("1. Download Video")
	cfmt.Cyan().Println("2. Download Audio")
	cfmt.Cyan().Println("3. Open in Browser")

	cfmt.Yellow().Print("==> ")
	fmt.Scan(&choice)

	renameExtension := ""

	switch choice {
	case 1:
		renameExtension = ".mp4"
		err = ytdl.DownloadVideo(chosenResult.Id)
	case 2:
		renameExtension = ".mp3"
		err = ytdl.DownloadAudio(chosenResult.Id)
	case 3:
		err = openUrlInBrowser(chosenResult.Url)
	default:
		err = errors.New("INVALID CHOICE!!!!!")
	}

	if err != nil {
		cfmt.Red().Bold().Println("An error has occurred ", err)
		os.Exit(1)
	}

	if renameExtension != "" {
		err = os.Rename(chosenResult.Id+renameExtension, chosenResult.Title+renameExtension)
		if err != nil {
			cfmt.Red().Bold().Println("Error renaming the downloaded video ", err)
			os.Exit(1)
		}
	}
}

func openUrlInBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Run()
	case "windows":
		err = exec.Command("start", url).Run()
	default:
		return errors.New("unsupported operating system")
	}
	if err != nil {
		return err
	}

	return nil
}

func resultItem(result ytscrape.VideoResult) string {
	builder := new(strings.Builder)
	builder.WriteString(cfmt.Green().Bold().Sprint(result.Title))
	builder.WriteByte(' ')
	builder.WriteString(cfmt.Magenta().Sprint(result.Duration.String()))
	builder.WriteString("\n\t")
	builder.WriteString(cfmt.Bold().Sprint("Uploaded By: ") + result.Uploader.Title)

	return builder.String()
}
