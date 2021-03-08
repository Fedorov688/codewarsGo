package splitStrings

import "testing"

func compareSlice(sliceA []string, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for key, value := range sliceB {
		if value != sliceA[key] {
			return false
		}
	}
	return true
}

func TestSolution(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{""}, []string{}},
		{"2", args{"a"}, []string{"a_"}},
		{"3", args{"abc"}, []string{"ab", "c_"}},
		{"4", args{"dawsd"}, []string{"da", "ws", "d_"}},
		{"5", args{"awsaws"}, []string{"aw", "sa", "ws"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.text); !compareSlice(got, tt.want) {
				t.Errorf("FindUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
