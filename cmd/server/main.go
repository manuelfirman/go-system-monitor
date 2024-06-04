package main

import (
	"github.com/manuelfirman/go-system-monitor/internal/application"
	"github.com/manuelfirman/go-system-monitor/internal/server"
)

func main() {
	server := server.NewServer()
	application := application.NewApplicationDefault(server)

	go application.Run()

	err := application.Listen()
	if err != nil {
		panic(err)
	}
}
