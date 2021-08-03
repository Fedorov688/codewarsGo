package sumOfIntInStr

import (
	"testing"
)

func TestSumOfIntegersInString(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"12.4"}, 16},
		{"2", args{"h3ll0w0rld"}, 3},
		{"3", args{"2 + 3 = "}, 5},
		{"4", args{"Our company made approximately 1 million in gross revenue last quarter."}, 1},
		{"5", args{"The Great Depression lasted from 1929 to 1939."}, 3868},
		{"6", args{"Dogs are our best friends."}, 0},
		{"7", args{"The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog"}, 3635},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfIntegersInString(tt.args.text); !(got == tt.want) {
				t.Errorf("FindUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
