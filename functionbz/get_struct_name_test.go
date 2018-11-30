package functionbz

import "testing"

type Bigzhu struct {
	big string
}

func toI(i interface{}) interface{} {
	return i
}

func TestGetStructName(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"base", args{Bigzhu{}}, "Bigzhu"},
		{"test pint", args{&Bigzhu{}}, "Bigzhu"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStructName(tt.args.s); got != tt.want {
				t.Errorf("GetStructName() = %v, want %v", got, tt.want)
			}
		})
	}
}
