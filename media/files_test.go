package media

import (
	"testing"
)

func TestIsMediaFile(t *testing.T) {
	// Supported media ".gif", ".jpeg", ".png"
	tests := []struct {
		name string
		path string
		want bool
	}{
		{name: "jpg", path: "c:\\temp\\otherdir\\hang.jpg", want: true},
		{name: "JPG in caps", path: "c:\\temp\\thingy.JPG", want: true},
		{name: "png", path: "c:\\temp\\thingy.png", want: true},
		{name: "gif", path: "c:\\temp\\thingy.gif", want: true},
		{name: "xps", path: "c:\\temp\\thingy.xps", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMediaFile(tt.path); got != tt.want {
				t.Errorf("IsMediaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
