package log_test

import "testing"

func RunBenchmark(b *testing.B, f func(int) int) {
	b.ReportAllocs()
	values := []int{0, 10, 20, 100, 1000}
	for i := 0; i < b.N; i++ {
		for _, value := range values {
			if g := f(value); g <= 0 {
				b.Fatalf("Unexpected got g(%d) <= 0", g)
			}
		}
	}
}
