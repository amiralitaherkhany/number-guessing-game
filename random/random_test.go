package random

import "testing"

func BenchmarkGenerateNumber(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		GenerateNumber(1, 100)
	}
}
