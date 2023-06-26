package main

import "attendance-record/shared/util"

func main() {
	util.CheckIsAppRunning()

	c := initClient()
	c.Run()
}
