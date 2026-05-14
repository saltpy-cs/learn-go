package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fname := os.Args[1]
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}
