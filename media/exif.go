package media

import (
	"os"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func extractMetaData(file *os.File) (metaData map[string]string) {
	fieldExtractionList := [...]exif.FieldName{exif.ImageDescription, exif.UserComment, exif.Artist, exif.MakerNote}
	metaData = make(map[string]string)
	decodedData, err := exif.Decode(file)

	// handle errors from decoding
	if err != nil {
		if strings.HasPrefix(err.Error(), "EOF") {
			// no meta data was found
			ConsoleLogger.Info("No meta data was found in", file.Name())
		} else {
			ConsoleLogger.Error("extractMetaData: ", err)
			ConsoleLogger.Trace("extractMetaData: Aborting")
		}

		return
	}

	j, _ := decodedData.MarshalJSON()
	ConsoleLogger.Info("Json:", string(j))

	for _, field := range fieldExtractionList {
		fieldValue := extractField(field, decodedData)
		if len(fieldValue) > 0 {
			metaData[string(field)] = strings.TrimSpace(fieldValue)
		}
	}

	return
}

func extractField(field exif.FieldName, decodedExif *exif.Exif) (returnValue string) {
	tagValue, error := decodedExif.Get(field)

	if error != nil {
		ConsoleLogger.Info("extractField: ", error)
	}
	//tagValue, _ := decodedExif.Get(exif.UserComment)
	if tagValue != nil {
		returnValue = strings.TrimSpace(tagValue.String())
	}
	return
}
