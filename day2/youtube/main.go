package main

import (
	"fmt"
	"os"

	"github.com/mbaraa/ytscrape"
)

func main() {
	results, err := ytscrape.Search("Luke Smith")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, result := range results {
		fmt.Println("Title", result.Title, "Channel", result.Uploader.Title, "Duration", result.Duration)
	}
}
