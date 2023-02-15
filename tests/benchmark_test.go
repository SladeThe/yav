package tests

import (
	"testing"

	"github.com/SladeThe/yav"
)

func BenchmarkYAV(b *testing.B) {
	account := ValidAccount()

	if err := account.Validate(); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = account.Validate()
	}
}

func BenchmarkOzzo(b *testing.B) {
	account := ValidAccount()

	if err := account.OzzoValidate(); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = account.OzzoValidate()
	}
}

func BenchmarkPlayground(b *testing.B) {
	v := yav.NewPlayground()
	account := ValidAccount()

	if err := v.Validate(account); err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = v.Validator.Struct(account)
	}
}

func BenchmarkYAVParallel(b *testing.B) {
	account := ValidAccount()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = account.Validate()
		}
	})
}

func BenchmarkOzzoParallel(b *testing.B) {
	account := ValidAccount()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = account.OzzoValidate()
		}
	})
}

func BenchmarkPlaygroundParallel(b *testing.B) {
	v := yav.NewPlayground()
	account := ValidAccount()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = v.Validator.Struct(account)
		}
	})
}
