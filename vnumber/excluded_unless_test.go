// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package vnumber

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestExcludedUnlessInt(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessInt(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessInt8(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           int8
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessInt8(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessInt16(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           int16
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessInt16(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessInt32(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           int32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessInt32(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessInt64(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           int64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessInt64(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessUint(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           uint
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessUint(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessUint8(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           uint8
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessUint8(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessUint16(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           uint16
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessUint16(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessUint32(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           uint32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessUint32(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessUint64(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           uint64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessUint64(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessFloat32(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           float32
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessFloat32(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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

func TestExcludedUnlessFloat64(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           float64
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnlessFloat64(a.conditionString, a.condition)(a.name, a.value)
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
			value:           1,
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
			value:           1,
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
			value:           0,
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