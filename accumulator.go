package yav

type accumulateFunc func(enabled, isEmpty bool) bool

var (
	requiredWithAnyAccumulateFunc    accumulateFunc = func(enabled, isEmpty bool) bool { return enabled || !isEmpty }
	requiredWithoutAnyAccumulateFunc accumulateFunc = func(enabled, isEmpty bool) bool { return enabled || isEmpty }
	requiredWithAllAccumulateFunc    accumulateFunc = func(enabled, isEmpty bool) bool { return enabled && !isEmpty }
	requiredWithoutAllAccumulateFunc accumulateFunc = func(enabled, isEmpty bool) bool { return enabled && isEmpty }
)

// TODO different type of Accumulator for each case (accumulators package), remove accumulateFunc from the structure

type Accumulator struct {
	accumulateFunc

	enabled bool
}

func (a Accumulator) String(value string) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == "")
	return a
}

func (a Accumulator) Bytes(value []byte) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, len(value) == 0)
	return a
}

func (a Accumulator) Bool(value bool) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, !value)
	return a
}

func (a Accumulator) Int(value int) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Int8(value int8) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Int16(value int16) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Int32(value int32) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Int64(value int64) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Uint(value int) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Uint8(value uint8) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Uint16(value uint16) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Uint32(value uint32) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) Uint64(value uint64) Accumulator {
	a.enabled = a.accumulateFunc(a.enabled, value == 0)
	return a
}

func (a Accumulator) IsEnabled() bool {
	return a.enabled
}

func RequiredWithAny() Accumulator {
	return Accumulator{
		accumulateFunc: requiredWithAnyAccumulateFunc,
		enabled:        false,
	}
}

func RequiredWithoutAny() Accumulator {
	return Accumulator{
		accumulateFunc: requiredWithoutAnyAccumulateFunc,
		enabled:        false,
	}
}

func RequiredWithAll() Accumulator {
	return Accumulator{
		accumulateFunc: requiredWithAllAccumulateFunc,
		enabled:        true,
	}
}

func RequiredWithoutAll() Accumulator {
	return Accumulator{
		accumulateFunc: requiredWithoutAllAccumulateFunc,
		enabled:        true,
	}
}
