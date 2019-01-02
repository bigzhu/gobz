package tokenize

import (
	"reflect"
	"testing"
)

func TestTextToWords(t *testing.T) {
	type args struct {
		text string
	}
	arg := args{text: `
	bigzhu
	bigzhu2
	`}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"base", arg, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TextToWords(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
