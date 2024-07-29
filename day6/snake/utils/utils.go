package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

func MoveCursorToBeginning() {
	fmt.Println("\033[1D")
}

func GetTermSize() (x, y int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	outSpl := strings.Split(string(out), " ")

	y, err = strconv.Atoi(outSpl[0])
	if err != nil {
		return 0, 0, err
	}

	x, err = strconv.Atoi(outSpl[1][:strings.Index(outSpl[1], "\n")])
	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}
