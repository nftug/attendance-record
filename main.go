package main

import (
	"client"
	"client/model"
)

func main() {
	api := model.NewApi()
	client.ShowAndRun(api)
}
