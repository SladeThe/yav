package yav

import (
	"errors"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsAppend(t *testing.T) {
	type args struct {
		err  Errors
		errs []error
	}

	type want struct {
		errs []Errors
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			errs := make([]Errors, len(a.errs))
			for i, err := range a.errs {
				errs[i].Append(a.err)
				errs[i].Append(err)
			}

			assert.Truef(t, reflect.DeepEqual(w.errs, errs), "invalid errors: want = %v, got = %v", w.errs, errs)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "validation errors with spare capacity",
		args: args{
			err: Errors{
				Validation: []Error{ErrRequired("a"), {}}[:1],
			},
			errs: []error{
				Errors{Validation: []Error{ErrRequired("b")}},
				ErrRequired("c"),
			},
		},
		want: want{
			errs: []Errors{
				{Validation: []Error{ErrRequired("a"), ErrRequired("b")}},
				{Validation: []Error{ErrRequired("a"), ErrRequired("c")}},
			},
		},
	}, {
		name: "unknown errors with spare capacity",
		args: args{
			err: Errors{
				Unknown: []error{errors.New("a"), nil}[:1],
			},
			errs: []error{
				Errors{Unknown: []error{errors.New("b")}},
				errors.New("c"),
			},
		},
		want: want{
			errs: []Errors{
				{Unknown: []error{errors.New("a"), errors.New("b")}},
				{Unknown: []error{errors.New("a"), errors.New("c")}},
			},
		},
	}, {
		name: "empty unknown errors",
		args: args{
			err: Errors{
				Unknown:    []error{},
				Validation: []Error{ErrRequired("v")},
			},
			errs: []error{nil},
		},
		want: want{
			errs: []Errors{{Validation: []Error{ErrRequired("v")}}},
		},
	}, {
		name: "empty validation errors",
		args: args{
			err: Errors{
				Unknown:    []error{errors.New("unknown")},
				Validation: []Error{},
			},
			errs: []error{nil},
		},
		want: want{
			errs: []Errors{{Unknown: []error{errors.New("unknown")}}},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestErrorsAppendConverted(t *testing.T) {
	type args struct {
		errs        Errors
		err         error
		convertFunc func(Error) Error
	}

	type want struct {
		errs  Errors
		calls int
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			originalErr := a.err
			if typedErr, ok := originalErr.(Errors); ok {
				typedErr.Unknown = slices.Clone(typedErr.Unknown)
				typedErr.Validation = slices.Clone(typedErr.Validation)
				originalErr = typedErr
			}

			calls := 0
			convertFunc := a.convertFunc
			if convertFunc != nil {
				convertFunc = func(err Error) Error {
					calls++
					return a.convertFunc(err)
				}
			}

			a.errs.AppendConverted(a.err, convertFunc)
			assert.Truef(t, reflect.DeepEqual(w.errs, a.errs), "invalid errors: want = %v, got = %v", w.errs, a.errs)
			assert.Equalf(t, w.calls, calls, "invalid calls: want = %v, got = %v", w.calls, calls)
			assert.Truef(t, reflect.DeepEqual(originalErr, a.err), "input changed: want = %v, got = %v", originalErr, a.err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "nil error",
		args: args{
			errs: Errors{Validation: []Error{ErrRequired("existing")}},
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted")
			},
		},
		want: want{
			errs: Errors{Validation: []Error{ErrRequired("existing")}},
		},
	}, {
		name: "nil convertFunc",
		args: args{
			err: Errors{
				Unknown:    []error{errors.New("unknown")},
				Validation: []Error{ErrRequired("v")},
			},
		},
		want: want{
			errs: Errors{
				Unknown:    []error{errors.New("unknown")},
				Validation: []Error{ErrRequired("v")},
			},
		},
	}, {
		name: "unknown error",
		args: args{
			err: errors.New("unknown"),
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted")
			},
		},
		want: want{
			errs: Errors{Unknown: []error{errors.New("unknown")}},
		},
	}, {
		name: "aggregate without validation errors",
		args: args{
			err: Errors{
				Unknown:    []error{errors.New("unknown")},
				Validation: []Error{},
			},
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted")
			},
		},
		want: want{
			errs: Errors{Unknown: []error{errors.New("unknown")}},
		},
	}, {
		name: "validation error",
		args: args{
			err: Error{
				CheckName: "min",
				Parameter: "1",
				ValueName: "v",
				Value:     0,
			},
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted")
			},
		},
		want: want{
			errs: Errors{Validation: []Error{{
				CheckName: "min",
				Parameter: "1",
				ValueName: "converted",
				Value:     0,
			}}},
			calls: 1,
		},
	}, {
		name: "validation aggregate with empty unknown errors",
		args: args{
			err: Errors{
				Unknown:    []error{},
				Validation: []Error{ErrRequired("v")},
			},
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted")
			},
		},
		want: want{
			errs:  Errors{Validation: []Error{ErrRequired("converted")}},
			calls: 1,
		},
	}, {
		name: "mixed aggregate with destination capacity",
		args: args{
			errs: Errors{
				Unknown:    []error{errors.New("existing"), nil}[:1],
				Validation: []Error{ErrRequired("existing"), {}, {}}[:1],
			},
			err: Errors{
				Unknown: []error{errors.New("unknown")},
				Validation: []Error{{
					CheckName: "min",
					Parameter: "1",
					ValueName: "low",
					Value:     0,
				}, {
					CheckName: "max",
					Parameter: "10",
					ValueName: "high",
					Value:     11,
				}},
			},
			convertFunc: func(err Error) Error {
				return err.WithValueName("converted." + err.ValueName)
			},
		},
		want: want{
			errs: Errors{
				Unknown: []error{errors.New("existing"), errors.New("unknown")},
				Validation: []Error{ErrRequired("existing"), {
					CheckName: "min",
					Parameter: "1",
					ValueName: "converted.low",
					Value:     0,
				}, {
					CheckName: "max",
					Parameter: "10",
					ValueName: "converted.high",
					Value:     11,
				}},
			},
			calls: 2,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
