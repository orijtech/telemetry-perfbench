package log_test

import (
	"strings"
	"testing"
)

func A(a int) int {
	if a > 0 {
		_ = 10 * 12
	}
	return B("Called from A")
}

func B(b string) int {
	b = strings.ToUpper(b)
	if len(b) > 1024 {
		b = strings.ToLower(b)
	}
	return len(b)
}
func BenchmarkNoTracingNoMetricsNoLogging(b *testing.B) {
	RunBenchmark(b, A)
}
