package main

import (
	"github.com/aliaksandrrachko/historian/historical-events/pkg/wiregen"
)

func main() {
	wiregen.InitServer().Start()
}
