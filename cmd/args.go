package cmd

import "flag"

type Args struct {
	Server *bool
}

func ArgParse() Args {
	args := Args{
		Server: flag.Bool("server", false, "Start the Валенки server"),
	}
	flag.BoolVar(args.Server, "s", false, "Start the Валенки server (shorthand)")
	flag.Parse()
	return args
}
