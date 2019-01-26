package main

import (
	cli "./cli"
	"./media"
	logrus "github.com/sirupsen/logrus"
)

//Logger
var ConsoleLogger logrus.Logger

// Application Entry Point
func main() {

	path, consoleLogLevel := cli.ParseArgs()
	ConsoleLogger = cli.GetLogger(cli.GetLogLevelFromString(consoleLogLevel))
	media.ConsoleLogger = ConsoleLogger
	cli.ConsoleLogger = ConsoleLogger
	ConsoleLogger.Info("Path:", path, " Console Log Level:", consoleLogLevel)
	ConsoleLogger.Trace("Looking for files")
	// set up the channels
	foundFiles := make(chan string)

	go media.FindFilesAsync(path, foundFiles)
	// media.GetMediaMetaData(<-foundFiles, log)

	for f := range foundFiles {
		ConsoleLogger.Trace("Ready to process file", f)
		mediaMetaData := media.GetMediaMetaData(f)
		cli.DisplayFileMetaData(f, mediaMetaData)

	}

}
