package main

import "attendance-record/client"

func main() {
	client.CheckIfAppRunning()
	c := initClient()
	c.Run()
}
