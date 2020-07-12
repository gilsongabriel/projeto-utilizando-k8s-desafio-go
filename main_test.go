package main

import (
	"reflect"
	"testing"
)

func Test_greeting(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"Test One",
			args{
				msg: "one",
			},
			[]byte("<b>one</b>"),
		},
		{
			"Test Two",
			args{
				msg: "two",
			},
			[]byte("<b>two</b>"),
		},
		{
			"Test Code Rocks",
			args{
				msg: "Code.education Rocks!",
			},
			[]byte("<b>Code.education Rocks!</b>"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greeting(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("greeting() = %v, want %v", got, tt.want)
			}
		})
	}
}