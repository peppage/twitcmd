package cmd

import (
	"fmt"
	"os"
	"strconv"
)

func er(msg interface{}) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(-1)
}

func newFalse() *bool {
	b := false
	return &b
}

func newTrue() *bool {
	b := true
	return &b
}

func parseID(id string) int64 {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		er("That tweet id is not a number")
	}
	return i
}
