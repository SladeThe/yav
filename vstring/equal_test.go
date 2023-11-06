package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestEqual(t *testing.T) {
	type args struct {
		parameter string
		name      string
		value     string
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
		name: "equal",
		args: args{
			parameter: "s",
			name:      "v",
			value:     "s",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not equal",
		args: args{
			parameter: "s",
			name:      "v",
			value:     "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: "s",
				ValueName: "v",
				Value:     "",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestNotEqual(t *testing.T) {
	type args struct {
		parameter string
		name      string
		value     string
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
		name: "not equal",
		args: args{
			parameter: "s",
			name:      "v",
			value:     "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "equal",
		args: args{
			parameter: "s",
			name:      "v",
			value:     "s",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameNotEqual,
				Parameter: "s",
				ValueName: "v",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestOneOf(t *testing.T) {
	type args struct {
		parameters []string
		name       string
		value      string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := OneOf(a.parameters...)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "in list",
		args: args{
			parameters: []string{"s0", "s1", "s2"},
			name:       "v",
			value:      "s0",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not in list",
		args: args{
			parameters: []string{"s0", "s1", "s2"},
			name:       "v",
			value:      "s",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameOneOf,
				Parameter: "s0 s1 s2",
				ValueName: "v",
				Value:     "s",
			},
		},
	}, {
		name: "empty list",
		args: args{
			parameters: nil,
			name:       "v",
			value:      "s",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameOneOf,
				Parameter: "",
				ValueName: "v",
				Value:     "s",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
