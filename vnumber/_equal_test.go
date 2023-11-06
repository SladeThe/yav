package vnumber

import (
	"reflect"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func TestEqualElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := EqualElement(a.parameter)(a.name, a.value)
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
			parameter: 1,
			name:      "v",
			value:     1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not equal",
		args: args{
			parameter: 1,
			name:      "v",
			value:     0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameEqual,
				Parameter: "1",
				ValueName: "v",
				Value:     Element(0),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestNotEqualElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := NotEqualElement(a.parameter)(a.name, a.value)
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
			parameter: 1,
			name:      "v",
			value:     0,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "equal",
		args: args{
			parameter: 1,
			name:      "v",
			value:     1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameNotEqual,
				Parameter: "1",
				ValueName: "v",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestOneOfElement(t *testing.T) {
	type args struct {
		parameters []Element
		name       string
		value      Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := OneOfElement(a.parameters...)(a.name, a.value)
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
			parameters: []Element{1, 2, 3},
			name:       "v",
			value:      1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "not in list",
		args: args{
			parameters: []Element{1, 2, 3},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameOneOf,
				Parameter: "1 2 3",
				ValueName: "v",
				Value:     Element(0),
			},
		},
	}, {
		name: "empty list",
		args: args{
			parameters: nil,
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameOneOf,
				Parameter: "",
				ValueName: "v",
				Value:     Element(0),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
