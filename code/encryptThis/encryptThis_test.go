package encryptThis

import "testing"

func TestEncryptThis(t *testing.T) {

	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{""}, ""},
		{"2", args{"A wise old owl lived in an oak"}, "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"},
		{"3", args{"The more he saw the less he spoke"}, "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"},
		{"4", args{"The less he spoke the more he heard"}, "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"},
		{"5", args{"Why can we not all be like that wise old bird"}, "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"},
		{"6", args{"Thank you Piotr for all your help"}, "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptThis(tt.args.text); got != tt.want {
				t.Errorf("EncryptThis() = %v, want %v", got, tt.want)
			}
		})
	}
}
