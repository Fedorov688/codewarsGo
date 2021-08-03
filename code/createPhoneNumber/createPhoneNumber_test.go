package createPhoneNumber

import "testing"

func TestCreatePhoneNumber(t *testing.T) {
	type args struct {
		text [10]uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, "(123) 456-7890"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPhoneNumber(tt.args.text); !(got == tt.want) {
				t.Errorf("FindUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
