package main

import (
	"PhotoMetadataHunter/cli"
	"PhotoMetadataHunter/logging"
	_ "PhotoMetadataHunter/logging"
	"PhotoMetadataHunter/media"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

var log logrus.Logger

// Application Entry Point
func main() {

	path, consoleLogLevel := cli.ParseArgs()
	log = logging.GetLogger(logging.GetLogLevelFromString(consoleLogLevel))
	log.Info("Path:", path, "Console Log Level:", consoleLogLevel)
	log.Trace("Looking for files")
	findFiles(path)

}

func findFiles(searchPath string) {
	var fileList []string
	filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if media.IsMediaFile(path) {
			log.Info("Found file:", path)
			fileList = append(fileList, path)
		}

		return nil
	})
}
