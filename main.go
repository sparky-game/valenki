package main

import (
	"flag"
	valenkiServer "github.com/sparky-game/valenki/pkg/server"
	"os"
)

func main() {
	serverFlag := flag.Bool("server", false, "Start the Валенки server")
	flag.BoolVar(serverFlag, "s", false, "Start the Валенки server (shorthand)")

	flag.Parse()

	if *serverFlag {
		valenkiServer.HTTP(6969)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
