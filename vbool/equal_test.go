package vbool

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestEqual(t *testing.T) {
	type args struct {
		parameter bool
		name      string
		value     bool
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Equal(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "equal true",
		args: args{
			parameter: true,
			name:      "v",
			value:     true,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not equal true",
		args: args{
			parameter: true,
			name:      "v",
			value:     false,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: "true",
				ValueName: "v",
			},
		},
	}, {
		name: "equal false",
		args: args{
			parameter: false,
			name:      "v",
			value:     false,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not equal false",
		args: args{
			parameter: false,
			name:      "v",
			value:     true,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: "false",
				ValueName: "v",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestNotEqual(t *testing.T) {
	type args struct {
		parameter bool
		name      string
		value     bool
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := NotEqual(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not equal true",
		args: args{
			parameter: true,
			name:      "v",
			value:     false,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "equal true",
		args: args{
			parameter: true,
			name:      "v",
			value:     true,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameNotEqual,
				Parameter: "true",
				ValueName: "v",
			},
		},
	}, {
		name: "not equal false",
		args: args{
			parameter: false,
			name:      "v",
			value:     true,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "equal false",
		args: args{
			parameter: false,
			name:      "v",
			value:     false,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameNotEqual,
				Parameter: "false",
				ValueName: "v",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
