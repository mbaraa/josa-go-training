package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

func GetTerminalSize() (width, height int, err error) {
	tputCols := exec.Command("tput", "cols")
	tputCols.Stdin = os.Stdin
	out, err := tputCols.Output()
	if err != nil {
		return 0, 0, err
	}
	cols, err := strconv.Atoi(string(out[:len(out)-1]))
	if err != nil {
		return 0, 0, err
	}

	tputLines := exec.Command("tput", "lines")
	tputLines.Stdin = os.Stdin
	out2, err := tputLines.Output()
	if err != nil {
		return 0, 0, err
	}
	lines, err := strconv.Atoi(string(out2[:len(out2)-1]))
	if err != nil {
		return 0, 0, err
	}

	return cols, lines, nil
}
