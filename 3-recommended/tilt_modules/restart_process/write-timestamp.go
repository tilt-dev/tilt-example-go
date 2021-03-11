package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// writes the current unix timestamp in nanos to a file
// used to trigger entr

func run() error {
	if len(os.Args) != 2 {
		fmt.Printf("usage: write-date <path>\n")
		os.Exit(1)
	}
	content := strconv.FormatInt(time.Now().UnixNano(), 10)
	err := ioutil.WriteFile(os.Args[1], []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}