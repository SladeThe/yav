package main

import (
	"fmt"
	"strings"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vpointer"
	"github.com/SladeThe/yav/vstring"
)

type Example struct {
	Optional *string
}

func (e Example) Validate() error {
	return yav.Chain(
		"Optional", e.Optional,
		vpointer.OmitEmpty[string], // Skip further validation, if pointer is nil.
		vpointer.Dereference[string](
			vstring.OmitEmpty, // Skip further validation, if string is empty.
			vstring.Max(100),
			vstring.Text,
		),
	)
}

func main() {
	valid := Example{
		Optional: func() *string {
			v := ""
			return &v
		}(),
	}

	invalid := Example{
		Optional: func() *string {
			v := strings.Repeat("1", 101)
			return &v
		}(),
	}

	fmt.Println(valid.Validate())
	fmt.Println(invalid.Validate())
}
