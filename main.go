package main

func main() {
	// get config
	c := config()

	// start server
	startServer(&c)
}
