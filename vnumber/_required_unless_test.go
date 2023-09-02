package vnumber

import (
	"reflect"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func TestRequiredUnlessElement(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := RequiredUnlessElement(a.conditionString, a.condition)(a.name, a.value)
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
			value:           0,
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
		name: "empty not required",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           0,
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
			value:           1,
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
