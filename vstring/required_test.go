package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestOmitEmpty(t *testing.T) {
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
			stop, err := OmitEmpty(a.name, a.value)
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
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			name:  "v",
			value: " ",
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

func TestRequired(t *testing.T) {
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
			stop, err := Required(a.name, a.value)
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
				CheckName: yav.CheckNameRequired,
				ValueName: "v",
			},
		},
	}, {
		name: "not empty",
		args: args{
			name:  "v",
			value: " ",
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

func TestRequiredIf(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := RequiredIf(a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredIf,
				Parameter: "a == a",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           " ",
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

func TestRequiredUnless(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := RequiredUnless(a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredUnless,
				Parameter: "a == b",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           " ",
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

func TestRequiredWithAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAny()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      " ",
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

func TestRequiredWithoutAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithoutAny()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      " ",
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

func TestRequiredWithAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAll()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      " ",
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

func TestRequiredWithoutAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithoutAll()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty required",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithoutAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      "",
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      " ",
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
