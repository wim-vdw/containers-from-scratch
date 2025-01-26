package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	checkLinux()
	run()
}

func checkLinux() {
	if runtime.GOOS != "linux" {
		panic("This program can only run on Linux")
	}
}

func run() {
	fmt.Printf("Running as pid: %d\n", os.Getpid())
	cmd := exec.Command("/bin/bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
