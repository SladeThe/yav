package vpointer

func OmitEmpty[T any](_ string, value *T) (stop bool, err error) {
	return value == nil, nil
}
