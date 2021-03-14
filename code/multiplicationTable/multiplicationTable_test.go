package multiplicationTable

import "testing"

func compareSlice(sliceA [][]int, sliceB [][]int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for key, value := range sliceB {
		if len(value) != len(sliceB[key]) {
			return false
		}
		for keyL2, valueL2 := range value {
			if valueL2 != sliceA[key][keyL2] {
				return false
			}
		}
	}
	return true
}

func TestMultiplicationTable(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{1}, [][]int{{1}}},
		{"2", args{2}, [][]int{{1, 2}, {2, 4}}},
		{"3", args{3}, [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplicationTable(tt.args.size); !compareSlice(got, tt.want) {
				t.Errorf("toNato() = %v, want %v", got, tt.want)
			}
		})
	}
}
