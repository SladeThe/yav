package common

import (
	"testing"
)

var (
	benchmarkTitles = []string{
		"Whispers in the Shadows",
		"The Silent Witness",
		"Love's Surrender",
		"A Love Like Yours",
		"The Mind Games",
		"Silent Whispers",
		"A Dark Obsession",
		"The Thread of Destiny",
		"Gone Without a Trace",
		"The Courage to Grow",
	}
)

func BenchmarkIsTitle(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for _, title := range benchmarkTitles {
			if !IsTitle(title) {
				b.Fatalf("not a title: %s", title)
			}
		}
	}
}

func BenchmarkIsTitleParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, title := range benchmarkTitles {
				if !IsTitle(title) {
					b.Fatalf("not a title: %s", title)
				}
			}
		}
	})
}
