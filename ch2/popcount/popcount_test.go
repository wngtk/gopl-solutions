package popcount_test

import (
	"testing"

	"github.com/wngtk/gopl-solutions/ch2/popcount"
)

func BenchmarkPopCountByAccumulate(b *testing.B) {
	for b.Loop() {
		popcount.PopCountByAccumulate(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for b.Loop() {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitCount(b *testing.B) {
	for b.Loop() {
		popcount.PopCountByBitCount(0x1234567890ABCDEF)
	}
}
