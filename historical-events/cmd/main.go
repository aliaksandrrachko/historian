package main

import (
	"github.com/aliaksandrrachko/historian/historical-events/pkg/server"
)

func main() {
	server := server.NewInstance()
	server.Start()
}
