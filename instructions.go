package main

import (
	"fmt"
	_ "fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> cmd args

func main() {

	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what??")
	}
}

func run() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getppid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"},os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
