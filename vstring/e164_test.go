package vstring

import (
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestE164(t *testing.T) {
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
			stop, err := E164(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid short",
		args: args{
			name:  "phone",
			value: "+1234567",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "valid long",
		args: args{
			name:  "phone",
			value: "+123456789101112",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid zero",
		args: args{
			name:  "phone",
			value: "+0234567891011",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "e164",
				ValueName: "phone",
				Value:     "+0234567891011",
			},
		},
	}, {
		name: "invalid short",
		args: args{
			name:  "phone",
			value: "+123456",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "e164",
				ValueName: "phone",
				Value:     "+123456",
			},
		},
	}, {
		name: "invalid long",
		args: args{
			name:  "phone",
			value: "+1234567891011121",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "e164",
				ValueName: "phone",
				Value:     "+1234567891011121",
			},
		},
	}, {
		name: "no plus sign",
		args: args{
			name:  "phone",
			value: "1234567891011",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "e164",
				ValueName: "phone",
				Value:     "1234567891011",
			},
		},
	}, {
		name: "invalid character",
		args: args{
			name:  "phone",
			value: "+123456789101a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "e164",
				ValueName: "phone",
				Value:     "+123456789101a",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkE164(b *testing.B) {
	var pp []string

	for i := 1; i <= 9; i++ {
		for j := 6; j <= 14; j++ {
			pp = append(pp, "+"+strconv.Itoa(i)+strings.Repeat("1", j))
		}
	}

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		for _, p := range pp {
			stop, err = E164("phone", p)
		}
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}
