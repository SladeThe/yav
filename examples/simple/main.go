package main

import (
	"fmt"
	"time"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vbool"
	"github.com/SladeThe/yav/vnumber"
	"github.com/SladeThe/yav/vstring"
	"github.com/SladeThe/yav/vtime"
)

type Example struct {
	Required      bool
	NonZero       int
	Positive      int16
	Email         string
	InaccurateNow time.Time
}

func (e Example) Validate() error {
	return yav.Join(
		yav.Chain(
			"Required", e.Required,
			vbool.Required,
		),
		yav.Chain(
			"NonZero", e.NonZero,
			vnumber.Required[int],
		),
		yav.Chain(
			"Positive", e.Positive,
			vnumber.GreaterThanInt16(0),
		),
		yav.Chain(
			"Email", e.Email,
			vstring.Required,
			vstring.Between(6, 100),
			vstring.Email,
		),
		yav.Chain(
			"InaccurateNow", e.InaccurateNow,
			vtime.Required,
			vtime.Between(time.Now().Add(-time.Second), time.Now().Add(time.Second)),
		),
	)
}

func main() {
	valid := Example{
		Required:      true,
		NonZero:       -1,
		Positive:      1,
		Email:         "example@example.org",
		InaccurateNow: time.Now(),
	}

	invalid := Example{
		Required:      false,
		NonZero:       0,
		Positive:      -1,
		Email:         "123",
		InaccurateNow: time.Now().Add(time.Minute),
	}

	fmt.Println(valid.Validate())
	fmt.Println(invalid.Validate())
}
