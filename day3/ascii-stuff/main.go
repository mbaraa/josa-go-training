package main

import (
	"fmt"
	"time"
)

func main() {
	// use \n for new lines
	// or \r\n on windows.
	fmt.Print("Hello hello\nNext line hehe!\n")

	// use \t to place a tab (4 spaces on most systems)
	// but tabs sizes can be of changed from the editor or text viewer's settings.
	fmt.Println("Hello\t\t\t\tfrom the other side")

	// Colors, well colors are a bit tricky,
	// each character modifier consists of three parts
	// style, foreground color, and background color.
	// and the final form has this structure \033[OPTION;FOREGROUND;BACKGROUNDm
	// e.g \033[1;35;47m
	// this represents a bold text (1) with megenta foreground (35) and white background (47)
	fmt.Println("\033[1;35;47mhello world")
	// after applying a text modifier, you need to reset it, so it's only reflected on the desired text.
	// use \0330m to reset the color
	fmt.Print("\033[0mnot so magenta anymore eh?")
	fmt.Println("\033[32m I'm green hehe\033[0m\033[34m I'm blue NOW!")
	fmt.Println("\033[0mAnd I'm normal again!")

	// Position control
	// like colors start with the escape sequence \033[ and use one of the known ASCII sequences,
	// full controls here https://en.wikipedia.org/wiki/ANSI_escape_code#Fe_Escape_sequences
	// for example \033[H\033[2J clears the screen and moves the cursor to the 0,0 coordinate
	// fmt.Println("\033[H\033[2J")

	spiny := "/-\\|"
	ticker := time.NewTicker(time.Second / 8)
	i := 0
	fmt.Println("Loading")
	for range ticker.C {
		fmt.Printf("\033[1D%c", spiny[i])
		i = (i + 1) % len(spiny)
	}
}
