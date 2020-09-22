package count

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_words(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"simple text",
			args{strings.NewReader(`foo foo bar, abc
			abc, 'foo', bar. and something more`)},
			map[string]int{
				"abc": 2,
				"foo": 3,
				"bar": 2,
				"and": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := words(tt.args.r)
			for wantWord, wantCount := range tt.want {
				if got[wantWord] != wantCount {
					t.Errorf("Words(): %s=%d, want %d",
						wantWord, got[wantWord], wantCount)
				}
			}
		})
	}
}

func TestWords(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []Word
	}{
		{
			"simple text",
			args{strings.NewReader(`abc, abc def. foo, 'bar', bar!
			abc, def, bar bar!`)},
			[]Word{Word{"bar", 4},
				Word{"abc", 3},
				Word{"def", 2},
				Word{"foo", 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Words(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Words() = %v, want %v", got, tt.want)
			}
		})
	}
}
