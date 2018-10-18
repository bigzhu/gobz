package websocketbz

import "testing"

func TestBroadcastByKey(t *testing.T) {
	type args struct {
		key string
		msg string
	}

	arg := args{key: "test", msg: "test_msg"}

	tests := []struct {
		name string
		args args
	}{
		{"base", arg},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BroadcastByKey(tt.args.key, tt.args.msg)
		})
	}
}
