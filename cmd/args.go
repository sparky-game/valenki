package cmd

import "flag"

type Args struct {
	Server *bool
}

func ArgParse() Args {
	args := Args{
		Server: flag.Bool("s", false, "Start the Валенки server"),
	}
	flag.Parse()
	return args
}
