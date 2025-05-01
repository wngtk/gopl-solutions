package popcount_test

import (
	"testing"

	mypopcount "github.com/wngtk/gopl-solutions/ch2/popcount"
	"gopl.io/ch2/popcount"
)

func BenchmarkPopCountByAccumulate(b *testing.B) {
	for b.Loop() {
		mypopcount.PopCountByAccumulate(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for b.Loop() {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByBitCount(b *testing.B) {
	for b.Loop() {
		mypopcount.PopCountByBitCount(0x1234567890ABCDEF)
	}
}
