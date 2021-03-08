package toNato

import "testing"

func TestToNato(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"If you can read"}, "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"},
		{"2", args{"Did not see that coming"}, "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"},
		{"3", args{"go for it!"}, "Golf Oscar Foxtrot Oscar Romeo India Tango !"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toNato(tt.args.text); got != tt.want {
				t.Errorf("toNato() = %v, want %v", got, tt.want)
			}
		})
	}
}
