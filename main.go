package main

import (
	"attendance-record/client"
	"attendance-record/client/model"
)

func main() {
	api := model.NewApi()
	client.ShowAndRun(api)
}
