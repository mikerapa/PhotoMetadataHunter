package media

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

//Logger
var ConsoleLogger logrus.Logger

// IsMediaFile Check to see if this file path is for a supported media file
func IsMediaFile(filePath string) bool {
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".jpg", ".gif", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

//FindFilesAsync finds media files and returns them to a chan
func FindFilesAsync(searchPath string, foundFilesChan chan string) {
	var fileList []string
	filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if IsMediaFile(path) {
			ConsoleLogger.Info("Found file:", path)
			fileList = append(fileList, path)
			foundFilesChan <- path
		}

		return nil
	})
	close(foundFilesChan)
}

// GetMediaMetaData - opens a file to extract metadata
func GetMediaMetaData(filePath string) (metaData map[string]string) {
	ConsoleLogger.Trace("GetMediaMetaData:", filePath)

	f, err := os.Open(filePath)
	if err != nil {
		ConsoleLogger.Error(err)
	} else {
		metaData = extractMetaData(f)
		ConsoleLogger.Info("Metadata:", metaData)
	}
	f.Close()
	return
}
