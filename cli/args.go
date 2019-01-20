package cli

import "flag"

func ParseArgs() (path string, consoleLogLevel string) {
	flag.StringVar(&path, "path", ".", "Path")
	flag.StringVar(&consoleLogLevel, "consoleLogLevel", "", "Log Level (error, warning, info, debug, trace)")
	flag.Parse()
	return
}
