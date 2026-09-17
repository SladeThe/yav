package yav

import (
	"errors"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNested(t *testing.T) {
	type args struct {
		names []string
		err   error
	}

	type want struct {
		errs []error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			originalErr := a.err
			if typedErr, ok := originalErr.(Errors); ok {
				typedErr.Unknown = slices.Clone(typedErr.Unknown)
				typedErr.Validation = slices.Clone(typedErr.Validation)
				originalErr = typedErr
			}
			errs := make([]error, len(a.names))
			for i, name := range a.names {
				errs[i] = Nested(name, a.err)
			}

			assert.Truef(t, reflect.DeepEqual(w.errs, errs), "invalid errors: want = %v, got = %v", w.errs, errs)
			assert.Truef(t, reflect.DeepEqual(originalErr, a.err), "input changed: want = %v, got = %v", originalErr, a.err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "invalid collection key",
		args: args{
			names: []string{"m"},
			err: Error{
				CheckName: "lowercase",
				Parameter: "",
				ValueName: "[{DEF}].code",
				Value:     "DEF",
			},
		},
		want: want{
			errs: []error{Error{
				CheckName: "lowercase",
				Parameter: "",
				ValueName: "m[{DEF}].code",
				Value:     "DEF",
			}},
		},
	}, {
		name: "reused aggregate",
		args: args{
			names: []string{"left", "right"},
			err: Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "min",
					Parameter: "1",
					ValueName: "x",
					Value:     0,
				}},
			},
		},
		want: want{
			errs: []error{Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "min",
					Parameter: "1",
					ValueName: "left.x",
					Value:     0,
				}},
			}, Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "min",
					Parameter: "1",
					ValueName: "right.x",
					Value:     0,
				}},
			}},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestNamedCheck(t *testing.T) {
	type args struct {
		names []string
		err   error
	}

	type want struct {
		errs []error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			originalErr := a.err
			if typedErr, ok := originalErr.(Errors); ok {
				typedErr.Unknown = slices.Clone(typedErr.Unknown)
				typedErr.Validation = slices.Clone(typedErr.Validation)
				originalErr = typedErr
			}
			errs := make([]error, len(a.names))
			for i, name := range a.names {
				errs[i] = NamedCheck(name, a.err)
			}

			assert.Truef(t, reflect.DeepEqual(w.errs, errs), "invalid errors: want = %v, got = %v", w.errs, errs)
			assert.Truef(t, reflect.DeepEqual(originalErr, a.err), "input changed: want = %v, got = %v", originalErr, a.err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "validation error",
		args: args{
			names: []string{"custom"},
			err: Error{
				CheckName: "min",
				Parameter: "1",
				ValueName: "x",
				Value:     0,
			},
		},
		want: want{
			errs: []error{Error{
				CheckName: "custom",
				Parameter: "",
				ValueName: "x",
				Value:     0,
			}},
		},
	}, {
		name: "reused aggregate",
		args: args{
			names: []string{"left", "right"},
			err: Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "min",
					Parameter: "1",
					ValueName: "x",
					Value:     0,
				}},
			},
		},
		want: want{
			errs: []error{Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "left",
					Parameter: "",
					ValueName: "x",
					Value:     0,
				}},
			}, Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "right",
					Parameter: "",
					ValueName: "x",
					Value:     0,
				}},
			}},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
