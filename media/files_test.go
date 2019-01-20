package media

import (
	"testing"
)

func TestIsMediaFile2(t *testing.T) {
	t.Error("fail")
}

func TestIsMediaFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMediaFile(tt.args.filePath); got != tt.want {
				t.Errorf("IsMediaFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
