package vtime

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestExcludedIf(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedIf(a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not empty excluded",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedIf,
				Parameter: "a == a",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           time.Time{},
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

func TestExcludedUnless(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnless(a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not empty excluded",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedUnless,
				Parameter: "a == b",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           time.Time{},
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

func TestExcludedWithAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithAny()

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
		name: "not empty excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      time.Time{},
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

func TestExcludedWithoutAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithoutAny()

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
		name: "not empty excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      time.Time{},
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

func TestExcludedWithAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithAll()

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
		name: "not empty excluded",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      time.Time{},
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

func TestExcludedWithoutAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithoutAll()

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
		name: "not empty excluded",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not empty not excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      time.Now(),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      time.Time{},
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
