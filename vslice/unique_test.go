package vslice

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestUnique(t *testing.T) {
	type args struct {
		name  string
		value []uint
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Unique(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "unique nil",
		args: args{
			name:  "",
			value: nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 0",
		args: args{
			name:  "",
			value: []uint{},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 1",
		args: args{
			name:  "",
			value: []uint{1},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 2",
		args: args{
			name:  "",
			value: []uint{1, 2},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "duplicate 2",
		args: args{
			name:  "s",
			value: []uint{1, 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{1, 1},
			},
		},
	}, {
		name: "duplicate 3",
		args: args{
			name:  "s",
			value: []uint{2, 1, 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{2, 1, 1},
			},
		},
	}, {
		name: "duplicate 4",
		args: args{
			name:  "s",
			value: []uint{3, 2, 1, 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{3, 2, 1, 1},
			},
		},
	}, {
		name: "duplicate 5",
		args: args{
			name:  "s",
			value: []uint{4, 3, 2, 1, 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{4, 3, 2, 1, 1},
			},
		},
	}, {
		name: "duplicate 6a",
		args: args{
			name:  "s",
			value: []uint{5, 4, 3, 2, 1, 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{5, 4, 3, 2, 1, 1},
			},
		},
	}, {
		name: "duplicate 6b",
		args: args{
			name:  "s",
			value: []uint{1, 1, 2, 3, 4, 5},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "s",
				Value:     []uint{1, 1, 2, 3, 4, 5},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkUnique(b *testing.B) {
	const maxLength = 16

	rnd := rand.New(rand.NewSource(1432164634643))

	var ss [][]uint64

	for length := 0; length <= maxLength; length++ {
		s := make([]uint64, 0, length)
		m := make(map[uint64]struct{}, length)

		for len(m) < length {
			v := rnd.Uint64()

			if _, ok := m[v]; ok {
				continue
			}

			s = append(s, v)
			m[v] = struct{}{}
		}

		ss = append(ss, s)
	}

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		for _, s := range ss {
			stop, err = Unique("name", s)
		}
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}
