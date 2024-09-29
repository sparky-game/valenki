package cmd

import (
	"flag"
	valenkiServer "github.com/sparky-game/valenki/pkg/server"
	"os"
)

func Start() {
  args := ArgParse()
	if *args.Server {
		valenkiServer.HTTP(6969)
	} else {
		flag.Usage()
		os.Exit(1)
	}
}
