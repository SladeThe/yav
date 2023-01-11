package vstruct

import (
	"github.com/SladeThe/yav"
)

type Validatable interface {
	Validate() error
}

func NestedValidate[T Validatable](name string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, yav.Nested(name, err)
}

func Validate[T Validatable](_ string, value T) (stop bool, err error) {
	err = value.Validate()
	return err != nil, err
}
