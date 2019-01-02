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
10 Years of Silence

Hayes started his research by examining successful composers. He analyzed thousands of musical pieces produced between the years of 1685 to 1900. The central question that drove his work was, “How long after one becomes interested in music is it that one becomes world class.

Eventually, Hayes developed a list of 500 pieces that were played frequently by symphonies around the world and were considered to be the “masterworks” in the field. These 500 popular pieces were created by a total of 76 composers.
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
				// t.Errorf("TextToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
