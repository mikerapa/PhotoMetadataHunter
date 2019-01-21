package main

import (
	cli "./cli"
	"./media"
	logrus "github.com/sirupsen/logrus"
)

var log logrus.Logger

// Application Entry Point
func main() {

	path, consoleLogLevel := cli.ParseArgs()
	log = cli.GetLogger(cli.GetLogLevelFromString(consoleLogLevel))
	log.Info("Path:", path, "Console Log Level:", consoleLogLevel)
	log.Trace("Looking for files")
	//findFiles(path)
	foundFiles := make(chan string)
	go media.FindFilesAsync(path, foundFiles, log)
	// media.GetMediaMetaData(<-foundFiles, log)

	for f := range foundFiles {
		log.Trace("Ready to process file", f)
		media.GetMediaMetaData(f, log)
	}

}
