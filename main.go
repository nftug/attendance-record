package main

import "attendance-record/shared/util"

func main() {
	util.CheckIfAppRunning()

	c := initClient()
	c.Run()
}
