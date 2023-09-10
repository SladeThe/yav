package yav

type Validatable interface {
	Validate() error
}

type Zeroer interface {
	IsZero() bool
}

type ValidateFunc[T any] func(name string, value T) (stop bool, err error)

func Next[T any](string, T) (stop bool, err error) {
	return false, nil
}

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

func Or2[T any](validateFunc1, validateFunc2 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc1(name, value); err == nil {
			return
		}

		return validateFunc2(name, value)
	}
}

func Or3[T any](validateFunc1, validateFunc2, validateFunc3 ValidateFunc[T]) ValidateFunc[T] {
	return func(name string, value T) (stop bool, err error) {
		if stop, err = validateFunc1(name, value); err == nil {
			return
		}

		if stop, err = validateFunc2(name, value); err == nil {
			return
		}

		return validateFunc3(name, value)
	}
}

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

func UnnamedValidate[T Validatable](_ string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, err
}

func NestedValidate[T Validatable](name string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, Nested(name, err)
}
