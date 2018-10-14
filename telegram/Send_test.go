package telegram

import (
	"testing"
)

func TestSend(t *testing.T) {
	type args struct {
		msg string
	}
	arg := args{msg: "这群要被毁掉了"}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base", arg, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
