package main

import (
	"github.com/reeechart/booql/book/server"
)

func main() {
	server := server.NewServer("", 5000)
	server.Run()
}
