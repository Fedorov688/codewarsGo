package findUniq

import "testing"

func TestFindUniq(t *testing.T) {
	type args struct {
		text []float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"1", args{[]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}}, float32(2)},
		{"2", args{[]float32{0, 0, 0.55, 0, 0}}, float32(0.55)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUniq(tt.args.text); got != tt.want {
				t.Errorf("FindUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
