package main

import (
	"fmt"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vslice"
	"github.com/SladeThe/yav/vstring"
)

type Example struct {
	Phones []string
}

func (e Example) Validate() error {
	return yav.Chain(
		"Phones", e.Phones,
		vslice.Required[[]string],
		vslice.Max[[]string](100),
		vslice.Items[[]string](
			vstring.Required,
			vstring.E164,
		),
	)
}

func main() {
	valid := Example{
		Phones: []string{"+1234567890"},
	}

	invalid := Example{
		Phones: []string{""},
	}

	fmt.Println(valid.Validate())
	fmt.Println(invalid.Validate())
}
