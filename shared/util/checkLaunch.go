package util

import (
	"log"
	"os"

	"github.com/mitchellh/go-ps"
)

func CheckIsAppRunning() {
	pid := os.Getpid()
	proc, _ := ps.FindProcess(pid)
	procName := proc.Executable()

	procs, _ := ps.Processes()
	for _, p := range procs {
		if p.Executable() == procName && p.Pid() != pid {
			log.Fatal("the app is already running")
		}
	}
}
