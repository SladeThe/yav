package yav

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNested(t *testing.T) {
	type args struct {
		name string
		err  error
	}

	type want struct {
		err error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			err := Nested(a.name, a.err)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "invalid collection key",
		args: args{
			name: "m",
			err: Error{
				CheckName: "lowercase",
				Parameter: "",
				ValueName: "[{DEF}].code",
				Value:     "DEF",
			},
		},
		want: want{
			err: Error{
				CheckName: "lowercase",
				Parameter: "",
				ValueName: "m[{DEF}].code",
				Value:     "DEF",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
