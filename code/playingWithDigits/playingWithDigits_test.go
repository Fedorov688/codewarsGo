package playingWithDigits

import "testing"

func TestPlayingWithDigits(t *testing.T) {
	type args struct {
		n int
		p int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{89, 1}, 1},
		{"2", args{92, 1}, -1},
		{"3", args{695, 2}, 2},
		{"4", args{46288, 3}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playingWithDigits(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("playingWithDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
