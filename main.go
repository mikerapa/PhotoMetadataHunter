package main

import (
	//"PhotoMetadataHunter/cli"
	//"PhotoMetadataHunter/logging"
	//_ "PhotoMetadataHunter/logging"
	//"PhotoMetadataHunter/media"

	"./cli"
	"./logging"
	"./media"
	logrus "github.com/sirupsen/logrus"
)

var log logrus.Logger

// Application Entry Point
func main() {

	path, consoleLogLevel := cli.ParseArgs()
	log = logging.GetLogger(logging.GetLogLevelFromString(consoleLogLevel))
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

// func findFiles(searchPath string) {
// 	var fileList []string
// 	filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
// 		if media.IsMediaFile(path) {
// 			log.Info("Found file:", path)
// 			fileList = append(fileList, path)
// 		}

// 		return nil
// 	})
// }
