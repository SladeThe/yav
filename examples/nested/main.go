package main

import (
	"fmt"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vbool"
)

type Nested struct {
	Required bool
}

func (n Nested) Validate() error {
	return yav.Chain(
		"Required", n.Required,
		vbool.Required,
	)
}

type Example struct {
	Nested

	Named Nested
}

func (e Example) Validate() error {
	return yav.Join2(
		e.Nested.Validate(),                     // Errors are returned as is.
		yav.Nested("Named", e.Named.Validate()), // Value names of returned errors are prepended with "Named.".
	)
}

func main() {
	valid := Example{
		Nested: Nested{Required: true},
		Named:  Nested{Required: true},
	}

	invalid := Example{}

	fmt.Println(valid.Validate())
	fmt.Println(invalid.Validate())
}
