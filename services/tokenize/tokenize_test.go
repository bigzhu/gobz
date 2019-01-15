package tokenize

import (
	"log"
	"testing"
)

func TestTextToWords(t *testing.T) {
	type args struct {
		text string
	}
	arg := args{text: `
	don't he's 
	“Follow your fuck!“
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
			got := TextToWords(tt.args.text)
			for _, i := range got {
				log.Printf("'%v'", i)
			}
		})
	}
}
