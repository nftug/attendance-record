package main

func main() {
	a := initApp()
	c := initClient(a)
	c.Run()
}
