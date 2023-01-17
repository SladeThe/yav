package accumulators

import (
	"github.com/SladeThe/yav"
)

type ProvideFunc[T any] func(fields string, enabled bool) yav.ValidateFunc[T]
