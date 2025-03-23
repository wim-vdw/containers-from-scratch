package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
)

func main() {
	must(checkLinux())
	must(checkRoot())
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	}
}

func checkLinux() error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("this program can only run on Linux")
	}
	return nil
}

func checkRoot() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("this program must be run as root")
	}
	return nil
}

func run() {
	fmt.Printf("Running %v as pid: %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as pid: %d\n", os.Args[2:], os.Getpid())
	cg()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/home/wim/alpinefs"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))
	must(cmd.Run())
	must(syscall.Unmount("/proc", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup"
	myCgroup := filepath.Join(cgroups, "container-from-scratch")
	must(os.MkdirAll(myCgroup, 0755))
	must(os.WriteFile(filepath.Join(myCgroup, "pids.max"), []byte("10"), 0700))
	must(os.WriteFile(filepath.Join(myCgroup, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
