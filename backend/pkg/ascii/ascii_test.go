package ascii_test

import (
	"testing"

	"github.com/p1ass/id/backend/pkg/ascii"
)

func TestIsASCII(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "all lower alphabets",
			str:  "hello",
			want: true,
		},
		{
			name: "all upper alphabets",
			str:  "HELLO",
			want: true,
		},
		{
			name: "combination of lower and upper alphabets",
			str:  "heLLO",
			want: true,
		},
		{
			name: "only numbers",
			str:  "1234",
			want: true,
		},
		{
			name: "combination of alphabets and numbers",
			str:  "hello1234",
			want: true,
		},
		{
			name: "only symbol",
			str:  "!@#$%^&*()_+-=",
			want: true,
		},
		{
			name: "include space",
			str:  "he llo",
			want: true,
		},
		{
			name: "all",
			str:  " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			want: true,
		},
		{
			name: "out of ASCII",
			str:  "こんにちは",
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := ascii.IsASCII(tt.str); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
