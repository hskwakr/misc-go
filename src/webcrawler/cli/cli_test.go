package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "case 1: Proper",
			in:   "webcrawler example.com http://www.foo.bar/index.html",
			want: 0,
		},
		{
			name: "case 2: Arguments error with no argument",
			in:   "webcrawler",
			want: 1,
		},
		{
			name: "case 2: Arguments error with 1 argument",
			in:   "webcrawler example.com",
			want: 1,
		},
	}

	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	app := &CLI{outStream, errStream}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			args := strings.Split(test.in, " ")

			status := app.Run(args)
			if status != test.want {
				t.Errorf("ExitStatus: %d, want: %d", status, test.want)
			}
		})
	}
}
