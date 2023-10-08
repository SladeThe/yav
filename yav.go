package yav

type Validatable interface {
	Validate() error
}

type Zeroer interface {
	IsZero() bool
}

// ValidateFunc usually represents a single validation check against the given named value.
type ValidateFunc[T any] func(name string, value T) (stop bool, err error)

// Next is a no-op ValidateFunc.
func Next[T any](string, T) (stop bool, err error) {
	return false, nil
}

// OmitEmpty skips further validation, when a generic comparable value is default for its type.
// Most of the validation packages contain specialized versions of this function, e.g. vstring.OmitEmpty, etc.
func OmitEmpty[T comparable](_ string, value T) (stop bool, err error) {
	var zero T
	return value == zero, nil
}

// Chain allows chaining validation funcs against a single struct field or value.
// If not nil, the result is always of Errors type.
func Chain[T any](name string, value T, validateFuncs ...ValidateFunc[T]) error {
	var yavErrs Errors

	for _, validateFunc := range validateFuncs {
		stop, err := validateFunc(name, value)
		yavErrs.Append(err)
		if stop {
			break
		}
	}

	return yavErrs.AsError()
}

// Join returns combined Errors or nil. It is useful to combine Chain results, while validating multiple values.
func Join(errs ...error) error {
	var yavErrs Errors

	for _, err := range errs {
		yavErrs.Append(err)
	}

	return yavErrs.AsError()
}

// Join2 exactly equals to Join with two arguments, but works faster.
func Join2(err0, err1 error) error {
	var yavErrs Errors

	yavErrs.Append(err0)
	yavErrs.Append(err1)

	return yavErrs.AsError()
}

// Join3 exactly equals to Join with three arguments, but works faster.
func Join3(err0, err1, err2 error) error {
	var yavErrs Errors

	yavErrs.Append(err0)
	yavErrs.Append(err1)
	yavErrs.Append(err2)

	return yavErrs.AsError()
}

// Or combines the given validation funcs into a new one, which iterates over and sequentially invokes the arguments.
// When any of the functions returns a nil error, its result is immediately returned.
// Otherwise, a non-nil error and stop flag of the last function are returned.
func Or[T any](validateFuncs ...ValidateFunc[T]) ValidateFunc[T] {
	if len(validateFuncs) == 1 {
		return validateFuncs[0]
	}

	return func(name string, value T) (stop bool, err error) {
		for _, validateFunc := range validateFuncs {
			if stop, err = validateFunc(name, value); err == nil {
				return
			}
		}

		return
	}
}

// Or2 exactly equals to Or with two arguments, but makes one less memory allocation.
func Or2[T any](validateFunc0, validateFunc1 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc0(name, value); err == nil {
			return
		}

		return validateFunc1(name, value)
	}
}

// Or3 exactly equals to Or with three arguments, but makes one less memory allocation.
func Or3[T any](validateFunc0, validateFunc1, validateFunc2 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc0(name, value); err == nil {
			return
		}

		if stop, err = validateFunc1(name, value); err == nil {
			return
		}

		return validateFunc2(name, value)
	}
}

// Nested processes errors of either Error or Errors type, prepending Error.ValueName with name argument.
// It returns unsupported and nil errors as is.
func Nested(name string, err error) error {
	if err == nil {
		return nil
	}

	switch typedErr := err.(type) {
	case Error:
		return nestedYAV(name, typedErr)
	case Errors:
		for i, yavErr := range typedErr.Validation {
			typedErr.Validation[i] = nestedYAV(name, yavErr)
		}

		return typedErr
	default:
		return err
	}
}

func nestedYAV(name string, yavErr Error) Error {
	if yavErr.ValueName == "" {
		yavErr.ValueName = name
	} else if yavErr.ValueName[0] == '[' {
		yavErr.ValueName = name + yavErr.ValueName
	} else {
		yavErr.ValueName = name + "." + yavErr.ValueName
	}

	return yavErr
}

// UnnamedValidate is a ValidateFunc that simply calls Validatable.Validate of the given value.
// It may be useful, while validating slice items or map entries.
func UnnamedValidate[T Validatable](_ string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, err
}

// NestedValidate basically equals to UnnamedValidate, but additionally calls Nested before returning the error.
func NestedValidate[T Validatable](name string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, Nested(name, err)
}
