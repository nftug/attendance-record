package util

import (
	"log"
	"os"

	"github.com/mitchellh/go-ps"
)

func CheckIfAppRunning() {
	pid := os.Getpid()
	curProc, _ := ps.FindProcess(pid)

	procs, _ := ps.Processes()
	for _, p := range procs {
		if p.Executable() == curProc.Executable() && p.Pid() != pid {
			log.Fatal("the app is already running")
		}
	}
}
