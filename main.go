package main

import (
	"fmt"
	"runtime"
)

func checkLinux() {
	if runtime.GOOS != "linux" {
		panic("This program can only run on Linux")
	}
}

func main() {
	checkLinux()
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
}
