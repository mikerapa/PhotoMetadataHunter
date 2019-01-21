package media

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

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
func FindFilesAsync(searchPath string, foundFilesChan chan string, consoleLogger logrus.Logger) {
	var fileList []string
	filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
		if IsMediaFile(path) {
			consoleLogger.Info("Found file:", path)
			fileList = append(fileList, path)
			foundFilesChan <- path
		}

		return nil
	})
	close(foundFilesChan)
}

// GetMediaMetaData - opens a file to extract metadata
func GetMediaMetaData(filePath string, consoleLogger logrus.Logger) (metaData map[string]string) {
	consoleLogger.Trace("GetMediaMetaData:", filePath)
	time.Sleep(100 * time.Millisecond) // TODO Remove this

	f, err := os.Open(filePath)
	if err != nil {
		consoleLogger.Error(err)
	} else {
		metaData = extractMetaData(f)
		consoleLogger.Info(metaData)
	}
	f.Close()
	return
}
