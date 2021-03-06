package cli

import (
	"testing"
)

func Test_trimValue(t *testing.T) {

	tests := []struct {
		name            string
		inputValue      string
		wantReturnValue string
	}{
		{name: "with tabs", inputValue: "		h	", wantReturnValue: "h"},
		{name: "leading spaces", inputValue: "  my value", wantReturnValue: "my value"},
		{name: "trailing spaces", inputValue: "my value  ", wantReturnValue: "my value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReturnValue := trimValue(tt.inputValue); gotReturnValue != tt.wantReturnValue {
				t.Errorf("trimValue() = %v, want %v", gotReturnValue, tt.wantReturnValue)
			}
		})
	}
}
