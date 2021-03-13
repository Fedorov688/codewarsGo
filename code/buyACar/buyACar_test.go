package buyACar

import "testing"

func TestToNato(t *testing.T) {
	type args struct {
		startPriceOld      int
		startPriceNew      int
		savingperMonth     int
		percentLossByMonth float64
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{"1", args{2000, 8000, 1000, 1.5}, [2]int{6, 766}},
		{"2", args{12000, 8000, 1000, 1.5}, [2]int{0, 4000}},
		{"3", args{8000, 12000, 500, 1.0}, [2]int{8, 597}},
		{"4", args{18000, 32000, 1500, 1.25}, [2]int{8, 332}},
		{"5", args{7500, 32000, 300, 1.55}, [2]int{25, 122}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NbMonths(tt.args.startPriceOld, tt.args.startPriceNew, tt.args.savingperMonth, tt.args.percentLossByMonth); got != tt.want {
				t.Errorf("NbMonths() = %v, want %v", got, tt.want)
			}
		})
	}
}
