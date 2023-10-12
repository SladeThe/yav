package vstring

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/SladeThe/yav"
)

const (
	base64RandomTestCount          = 4
	base64RandomDataMaxLengthBytes = 100

	base64BenchmarkDataLengthBytes = 1 << 10
)

var (
	base64BenchmarkBytes = func() []byte {
		bytes := make([]byte, base64BenchmarkDataLengthBytes)

		if _, errValue := rand.Read(bytes); errValue != nil {
			panic(errValue)
		}

		return bytes
	}()
)

func TestBase64(t *testing.T) {
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
			stop, err := Base64(a.name, a.value)
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
			name:  "data",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			name:  "data",
			value: "A?==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64",
				ValueName: "data",
				Value:     "A?==",
			},
		},
	}, {
		name: "wrong encoding 0",
		args: args{
			name:  "data",
			value: "AQ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64",
				ValueName: "data",
				Value:     "AQ",
			},
		},
	}, {
		name: "wrong encoding 1",
		args: args{
			name:  "data",
			value: "-_==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64",
				ValueName: "data",
				Value:     "-_==",
			},
		},
	}}

	for i := 0; i < base64RandomTestCount; i++ {
		length, errLength := rand.Int(rand.Reader, big.NewInt(base64RandomDataMaxLengthBytes))
		require.NoError(t, errLength)

		value := make([]byte, 1+length.Int64())
		_, errValue := rand.Read(value)
		require.NoError(t, errValue)

		tests = append(tests, struct {
			name string
			args args
			want want
		}{
			name: fmt.Sprintf("random valid %v", i),
			args: args{
				name:  "data",
				value: base64.StdEncoding.EncodeToString(value),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBase64Raw(t *testing.T) {
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
			stop, err := Base64Raw(a.name, a.value)
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
			name:  "data",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			name:  "data",
			value: "A?",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64raw",
				ValueName: "data",
				Value:     "A?",
			},
		},
	}, {
		name: "wrong encoding 0",
		args: args{
			name:  "data",
			value: "AQ==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64raw",
				ValueName: "data",
				Value:     "AQ==",
			},
		},
	}, {
		name: "wrong encoding 1",
		args: args{
			name:  "data",
			value: "-_",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64raw",
				ValueName: "data",
				Value:     "-_",
			},
		},
	}}

	for i := 0; i < base64RandomTestCount; i++ {
		length, errLength := rand.Int(rand.Reader, big.NewInt(base64RandomDataMaxLengthBytes))
		require.NoError(t, errLength)

		value := make([]byte, 1+length.Int64())
		_, errValue := rand.Read(value)
		require.NoError(t, errValue)

		tests = append(tests, struct {
			name string
			args args
			want want
		}{
			name: fmt.Sprintf("random valid %v", i),
			args: args{
				name:  "data",
				value: base64.RawStdEncoding.EncodeToString(value),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBase64URL(t *testing.T) {
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
			stop, err := Base64URL(a.name, a.value)
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
			name:  "data",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			name:  "data",
			value: "A?==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64url",
				ValueName: "data",
				Value:     "A?==",
			},
		},
	}, {
		name: "wrong encoding 0",
		args: args{
			name:  "data",
			value: "AQ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64url",
				ValueName: "data",
				Value:     "AQ",
			},
		},
	}, {
		name: "wrong encoding 1",
		args: args{
			name:  "data",
			value: "+/==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64url",
				ValueName: "data",
				Value:     "+/==",
			},
		},
	}}

	for i := 0; i < base64RandomTestCount; i++ {
		length, errLength := rand.Int(rand.Reader, big.NewInt(base64RandomDataMaxLengthBytes))
		require.NoError(t, errLength)

		value := make([]byte, 1+length.Int64())
		_, errValue := rand.Read(value)
		require.NoError(t, errValue)

		tests = append(tests, struct {
			name string
			args args
			want want
		}{
			name: fmt.Sprintf("random valid %v", i),
			args: args{
				name:  "data",
				value: base64.URLEncoding.EncodeToString(value),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBase64RawURL(t *testing.T) {
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
			stop, err := Base64RawURL(a.name, a.value)
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
			name:  "data",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			name:  "data",
			value: "A?",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64rawurl",
				ValueName: "data",
				Value:     "A?",
			},
		},
	}, {
		name: "wrong encoding 0",
		args: args{
			name:  "data",
			value: "AQ==",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64rawurl",
				ValueName: "data",
				Value:     "AQ==",
			},
		},
	}, {
		name: "wrong encoding 1",
		args: args{
			name:  "data",
			value: "+/",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "base64rawurl",
				ValueName: "data",
				Value:     "+/",
			},
		},
	}}

	for i := 0; i < base64RandomTestCount; i++ {
		length, errLength := rand.Int(rand.Reader, big.NewInt(base64RandomDataMaxLengthBytes))
		require.NoError(t, errLength)

		value := make([]byte, 1+length.Int64())
		_, errValue := rand.Read(value)
		require.NoError(t, errValue)

		tests = append(tests, struct {
			name string
			args args
			want want
		}{
			name: fmt.Sprintf("random valid %v", i),
			args: args{
				name:  "data",
				value: base64.RawURLEncoding.EncodeToString(value),
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkBase64(b *testing.B) {
	base64Value := base64.StdEncoding.EncodeToString(base64BenchmarkBytes)

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		stop, err = Base64("data", base64Value)
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}

func BenchmarkBase64Raw(b *testing.B) {
	base64Value := base64.RawStdEncoding.EncodeToString(base64BenchmarkBytes)

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		stop, err = Base64Raw("data", base64Value)
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}

func BenchmarkBase64URL(b *testing.B) {
	base64Value := base64.URLEncoding.EncodeToString(base64BenchmarkBytes)

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		stop, err = Base64URL("data", base64Value)
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}

func BenchmarkBase64RawURL(b *testing.B) {
	base64Value := base64.RawURLEncoding.EncodeToString(base64BenchmarkBytes)

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		stop, err = Base64RawURL("data", base64Value)
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}
