package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

// TODO other functions

func TestIsTitle(t *testing.T) {
	type args struct {
		name  string
		value string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Title(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "leading space",
		args: args{
			name:  "v",
			value: " Title",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     " Title",
			},
		},
	}, {
		name: "leading tab",
		args: args{
			name:  "v",
			value: "\tTitle",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "\tTitle",
			},
		},
	}, {
		name: "trailing space",
		args: args{
			name:  "v",
			value: "Title ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Title ",
			},
		},
	}, {
		name: "trailing tab",
		args: args{
			name:  "v",
			value: "Title\t",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Title\t",
			},
		},
	}, {
		name: "repeating spaces",
		args: args{
			name:  "v",
			value: "Two  words",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two  words",
			},
		},
	}, {
		name: "tab",
		args: args{
			name:  "v",
			value: "Two\twords",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two\twords",
			},
		},
	}, {
		name: "nbsp",
		args: args{
			name:  "v",
			value: "Two\u00A0words",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two\u00A0words",
			},
		},
	}, {
		name: "valid one word",
		args: args{
			name:  "v",
			value: "Title",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "valid two words",
		args: args{
			name:  "v",
			value: "Two words",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
