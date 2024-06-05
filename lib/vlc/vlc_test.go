package vlc

import "testing"

func Test_prepareText(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is GRISHA",
			want: "!my name is !g!r!i!s!h!a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}
