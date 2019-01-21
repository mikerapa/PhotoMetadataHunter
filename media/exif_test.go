package media

import (
	"os"
	"strings"
	"testing"
)

func Test_extractMetaData(t *testing.T) {
	f, err := os.Open("../test media/mt.jpg")
	if err != nil {
		t.Error("Unable to open image file")
	}
	returnedMetaData := extractMetaData(f)
	// Check for ImageDescription
	imageDescriptionValue, imageDescriptionFound := returnedMetaData["ImageDescription"]
	if !imageDescriptionFound {
		t.Error("Expecting a value for ImageDescription")
	}
	if !strings.Contains(string(imageDescriptionValue), "Title") {
		t.Error("Expected an ImageDescription with a specific value")
	}

	// User Comment should not exist
	_, userCommentFound := returnedMetaData["UserComment"]
	if userCommentFound {
		t.Error("Expecting no value for UserComment")
	}
	f.Close()
}
