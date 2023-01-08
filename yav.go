package yav

type ValidateFunc[T any] func(name string, value T) (stop bool, err error)

func Chain[T any](name string, value T, validateFuncs ...ValidateFunc[T]) error {
	for _, validateFunc := range validateFuncs {
		if stop, err := validateFunc(name, value); stop || err != nil {
			return err
		}
	}

	return nil
}

func Or[T any](name string, value T, validateFuncs ...ValidateFunc[T]) (stop bool, err error) {
	for _, validateFunc := range validateFuncs {
		if stop, err = validateFunc(name, value); stop || err == nil {
			return
		}
	}

	return
}
