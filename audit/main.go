package main

import "github.com/reeechart/booql/audit/server"

func main() {
	server := server.NewServer("", 5001)
	server.Run()
}
