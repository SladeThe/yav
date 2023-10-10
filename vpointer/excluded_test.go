package vpointer

import (
	"runtime"
	"testing"
	"time"

	"github.com/SladeThe/yav"
)

func BenchmarkExcludedAccumulators(b *testing.B) {
	var pointer *time.Time

	b.ReportAllocs()

	var err error

	for i := 0; i < b.N; i++ {
		err = yav.Join(
			yav.Chain(
				"pointer", pointer,
				ExcludedIf[time.Time]("", true),
				ExcludedUnless[time.Time]("", false),
				ExcludedWithAny[time.Time]().Bool(true).Names(""),
				ExcludedWithoutAny[time.Time]().Bool(false).Names(""),
				ExcludedWithAll[time.Time]().Bool(true).Names(""),
				ExcludedWithoutAll[time.Time]().Bool(false).Names(""),
			),
		)
	}

	runtime.KeepAlive(err)
}

func BenchmarkExcludedAccumulatorsParallel(b *testing.B) {
	var pointer *time.Time

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = yav.Join(
				yav.Chain(
					"pointer", pointer,
					ExcludedIf[time.Time]("", true),
					ExcludedUnless[time.Time]("", false),
					ExcludedWithAny[time.Time]().Bool(true).Names(""),
					ExcludedWithoutAny[time.Time]().Bool(false).Names(""),
					ExcludedWithAll[time.Time]().Bool(true).Names(""),
					ExcludedWithoutAll[time.Time]().Bool(false).Names(""),
				),
			)
		}
	})
}
