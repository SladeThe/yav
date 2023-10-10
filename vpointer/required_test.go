package vpointer

import (
	"runtime"
	"testing"
	"time"

	"github.com/SladeThe/yav"
)

func BenchmarkRequiredAccumulators(b *testing.B) {
	pointer := new(time.Time)

	b.ReportAllocs()

	var err error

	for i := 0; i < b.N; i++ {
		err = yav.Join(
			yav.Chain(
				"pointer", pointer,
				RequiredIf[time.Time]("", true),
				RequiredUnless[time.Time]("", false),
				RequiredWithAny[time.Time]().Bool(true).Names(""),
				RequiredWithoutAny[time.Time]().Bool(false).Names(""),
				RequiredWithAll[time.Time]().Bool(true).Names(""),
				RequiredWithoutAll[time.Time]().Bool(false).Names(""),
			),
		)
	}

	runtime.KeepAlive(err)
}

func BenchmarkRequiredAccumulatorsParallel(b *testing.B) {
	pointer := new(time.Time)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = yav.Join(
				yav.Chain(
					"pointer", pointer,
					RequiredIf[time.Time]("", true),
					RequiredUnless[time.Time]("", false),
					RequiredWithAny[time.Time]().Bool(true).Names(""),
					RequiredWithoutAny[time.Time]().Bool(false).Names(""),
					RequiredWithAll[time.Time]().Bool(true).Names(""),
					RequiredWithoutAll[time.Time]().Bool(false).Names(""),
				),
			)
		}
	})
}
