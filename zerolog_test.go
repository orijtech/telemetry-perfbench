package log_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func A_log_zerolog(logger zerolog.Logger, a int) int {
	if a > 0 {
		logger.Info().Msgf("a > 0 where a=%d", a)
		_ = 10 * 12
	}
	logger.Info().Msg("calling b")
	return B_log_zerolog(logger, "Called from A")
}

func B_log_zerolog(logger zerolog.Logger, b string) int {
	b = strings.ToUpper(b)
	logger.Info().Msgf("b uppercased, so lowercased where len_b=%d", len(b))
	if len(b) > 1024 {
		b = strings.ToLower(b)
		logger.Info().Msgf("b > 1024, so lowercased where b=%s", b)
	}
	return len(b)
}

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

func BenchmarkLoggingZerolog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard)
	b.ResetTimer()
	RunBenchmark(b, func(a int) int {
		return A_log_zerolog(logger, a)
	})
}
