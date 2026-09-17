package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestUUID(t *testing.T) {
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
			stop, err := UUID(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "lowercase",
		args: args{
			name:  "",
			value: "6a310c88-4698-4807-9578-f1f054a8b4ca",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "uppercase",
		args: args{
			name:  "id",
			value: "6A310C88-4698-4807-9578-F1F054A8B4CA",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "mixed case",
		args: args{
			name:  "id",
			value: "6a310C88-4698-4807-9578-f1f054a8B4ca",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid uppercase hex",
		args: args{
			name:  "id",
			value: "6A310C88-4698-4807-9578-G1F054A8B4CA",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uuid",
				ValueName: "id",
				Value:     "6A310C88-4698-4807-9578-G1F054A8B4CA",
			},
		},
	}, {
		name: "invalid length",
		args: args{
			name:  "id",
			value: "6a310c88-4698-4807-9578-f1f054a8b4ca1",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uuid",
				ValueName: "id",
				Value:     "6a310c88-4698-4807-9578-f1f054a8b4ca1",
			},
		},
	}, {
		name: "invalid part rune",
		args: args{
			name:  "id",
			value: "6a310c88-4698-4807-9578-g1f054a8b4ca",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uuid",
				ValueName: "id",
				Value:     "6a310c88-4698-4807-9578-g1f054a8b4ca",
			},
		},
	}, {
		name: "invalid separator rune",
		args: args{
			name:  "id",
			value: "6a310c88-4698-4807a9578-f1f054a8b4ca",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uuid",
				ValueName: "id",
				Value:     "6a310c88-4698-4807a9578-f1f054a8b4ca",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
