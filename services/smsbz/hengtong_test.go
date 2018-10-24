package smsbz

import "testing"

func TestSend(t *testing.T) {
	type args struct {
		phone string
		msg   string
	}
	arg := args{"18987128028", "这是一个测试,啊哈哈"}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", arg, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Send(tt.args.phone, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
