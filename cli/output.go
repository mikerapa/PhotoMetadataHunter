package cli

import (
	"fmt"
	"gopkg.in/gookit/color.v1"
)

func DisplayFileMetaData(fileName string, fileMetaData map[string]string) {
	fileColor := color.New(color.Black, color.BgYellow, color.Bold)
	fileColor.Print(fileName)
	fmt.Print("\t")
	firstProperty := true
	for k, v := range fileMetaData {
		if !firstProperty {
			color.White.Print(", ")
		}
		firstProperty = false
		color.Yellow.Print(k)
		color.White.Print("=")
		color.White.Print(v)
	}
	fmt.Print("\n")
}
