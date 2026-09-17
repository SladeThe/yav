package vstring

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestHostname(t *testing.T) {
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
			stop, err := Hostname(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "hostname",
		args: args{
			name:  "v",
			value: "host",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "short labels",
		args: args{
			name:  "v",
			value: "a.b",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "internal hyphen",
		args: args{
			name:  "v",
			value: "a-1.example",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname",
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "single character",
		args: args{
			name:  "v",
			value: "a",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "leading digit",
		args: args{
			name:  "v",
			value: "1host",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname",
				ValueName: "v",
				Value:     "1host",
			},
		},
	}, {
		name: "label ending with hyphen",
		args: args{
			name:  "v",
			value: "ab-.example",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname",
				ValueName: "v",
				Value:     "ab-.example",
			},
		},
	}, {
		name: "long label",
		args: args{
			name:  "v",
			value: strings.Repeat("a", 64),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname",
				ValueName: "v",
				Value:     strings.Repeat("a", 64),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestHostnameRFC1123(t *testing.T) {
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
			stop, err := HostnameRFC1123(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "single character",
		args: args{
			name:  "v",
			value: "a",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "leading digit",
		args: args{
			name:  "v",
			value: "1host.example",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "internal hyphen",
		args: args{
			name:  "v",
			value: "a-b.example",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "maximum label",
		args: args{
			name:  "v",
			value: strings.Repeat("a", 63),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "maximum name",
		args: args{
			name:  "v",
			value: strings.Repeat("a.", 126) + "a",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "leading dot",
		args: args{
			name:  "v",
			value: ".a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     ".a",
			},
		},
	}, {
		name: "empty label",
		args: args{
			name:  "v",
			value: "a..b",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a..b",
			},
		},
	}, {
		name: "trailing dot",
		args: args{
			name:  "v",
			value: "a.",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a.",
			},
		},
	}, {
		name: "leading hyphen",
		args: args{
			name:  "v",
			value: "-a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "-a",
			},
		},
	}, {
		name: "trailing hyphen",
		args: args{
			name:  "v",
			value: "a-",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a-",
			},
		},
	}, {
		name: "label ending with hyphen",
		args: args{
			name:  "v",
			value: "a-.b",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a-.b",
			},
		},
	}, {
		name: "label starting with hyphen",
		args: args{
			name:  "v",
			value: "a.-b",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a.-b",
			},
		},
	}, {
		name: "underscore",
		args: args{
			name:  "v",
			value: "a_b",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "a_b",
			},
		},
	}, {
		name: "non-ASCII",
		args: args{
			name:  "v",
			value: "é.example",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     "é.example",
			},
		},
	}, {
		name: "long label",
		args: args{
			name:  "v",
			value: strings.Repeat("a", 64),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     strings.Repeat("a", 64),
			},
		},
	}, {
		name: "long name",
		args: args{
			name:  "v",
			value: strings.Repeat("a.", 126) + "aa",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_rfc1123",
				ValueName: "v",
				Value:     strings.Repeat("a.", 126) + "aa",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestHostnamePort(t *testing.T) {
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
			stop, err := HostnamePort(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "hostname",
		args: args{
			name:  "v",
			value: "host:80",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty host",
		args: args{
			name:  "v",
			value: ":65535",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "leading digit",
		args: args{
			name:  "v",
			value: "1host:443",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "zero port",
		args: args{
			name:  "v",
			value: "host:0",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_port",
				ValueName: "v",
				Value:     "host:0",
			},
		},
	}, {
		name: "large port",
		args: args{
			name:  "v",
			value: "host:65536",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_port",
				ValueName: "v",
				Value:     "host:65536",
			},
		},
	}, {
		name: "invalid host",
		args: args{
			name:  "v",
			value: "foo-:80",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_port",
				ValueName: "v",
				Value:     "foo-:80",
			},
		},
	}, {
		name: "missing port",
		args: args{
			name:  "v",
			value: "host",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_port",
				ValueName: "v",
				Value:     "host",
			},
		},
	}, {
		name: "invalid port",
		args: args{
			name:  "v",
			value: "host:abc",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "hostname_port",
				ValueName: "v",
				Value:     "host:abc",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestFQDN(t *testing.T) {
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
			stop, err := FQDN(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "domain",
		args: args{
			name:  "v",
			value: "example.com",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "root dot",
		args: args{
			name:  "v",
			value: "example.com.",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "short labels",
		args: args{
			name:  "v",
			value: "a.b",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "punycode TLD",
		args: args{
			name:  "v",
			value: "example.xn--p1ai",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "maximum name with root dot",
		args: args{
			name:  "v",
			value: strings.Repeat("a.", 126) + "a.",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "single label",
		args: args{
			name:  "v",
			value: "localhost",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     "localhost",
			},
		},
	}, {
		name: "numeric TLD",
		args: args{
			name:  "v",
			value: "example.123",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     "example.123",
			},
		},
	}, {
		name: "trailing hyphen",
		args: args{
			name:  "v",
			value: "foo-.example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     "foo-.example.com",
			},
		},
	}, {
		name: "TLD ending with hyphen",
		args: args{
			name:  "v",
			value: "example.com-",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     "example.com-",
			},
		},
	}, {
		name: "empty label",
		args: args{
			name:  "v",
			value: "example..com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     "example..com",
			},
		},
	}, {
		name: "long name",
		args: args{
			name:  "v",
			value: strings.Repeat("a.", 126) + "aa.",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "fqdn",
				ValueName: "v",
				Value:     strings.Repeat("a.", 126) + "aa.",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
