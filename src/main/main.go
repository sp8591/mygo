package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"os"
	"log"
	"time"
	"path"
	"io/ioutil"
	"strconv"
)
func uts() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr= &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("dd")
}
func ipc() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("dd")
}

func mnt() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(i)
	}
	fmt.Println("dd")
}
const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func mem() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d\n", syscall.Getegid())
		cmd := exec.Command("sh", arg) 
	}
}
func main() {
	mnt()
}

