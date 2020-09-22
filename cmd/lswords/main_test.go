package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut []string
		wantErr string
	}{
		{
			"simple txt",
			args{strings.NewReader(`abc abc def abc`)},
			[]string{
				"3  75.00%  abc\n",
				"1  25.00%  def\n",
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			err := &bytes.Buffer{}
			run(tt.args.in, out, err)
			gotOut := out.String()
			for _, wOut := range tt.wantOut {
				if !strings.Contains(gotOut, wOut) {
					t.Errorf("run() should contain: \n%s\ngot:\n%s",
						wOut, gotOut)
				}
			}
			if gotErr := err.String(); gotErr != tt.wantErr {
				t.Errorf("run() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
