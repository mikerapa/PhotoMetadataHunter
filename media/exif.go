package media

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func extractMetaData(file *os.File) (metaData map[string]string) {
	fieldExtractionList := [...]exif.FieldName{exif.ImageDescription, exif.UserComment}
	metaData = make(map[string]string)
	decodedData, err := exif.Decode(file)
	if err != nil {
		// TODO add a way to log this error
	}
	for _, field := range fieldExtractionList {
		fieldValue := extractField(field, decodedData)
		if len(fieldValue) > 0 {
			metaData[string(field)] = fieldValue
		}
	}

	// metaData["UserComment"] = extractField(exif.UserComment, decodedData)
	// metaData["ImageDescription"] = extractField(exif.ImageDescription, decodedData)
	return
}

func extractField(field exif.FieldName, decodedExif *exif.Exif) (returnValue string) {
	tagValue, _ := decodedExif.Get(field)
	//tagValue, _ := decodedExif.Get(exif.UserComment)
	if tagValue != nil {
		returnValue = tagValue.String()
	}
	return
}
