package rgbToHex

import "testing"

func TestRgb(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{0, 0, 0}, "000000"},
		{"2", args{0, 0, -20}, "000000"},
		{"3", args{300, 255, 255}, "FFFFFF"},
		{"4", args{173, 255, 47}, "ADFF2F"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rgb(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("rgb() = %v, want %v", got, tt.want)
			}
		})
	}
}
