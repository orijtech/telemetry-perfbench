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

func BenchmarkLoggingZerolog(b *testing.B) {
	logger := zerolog.New(ioutil.Discard)
	b.ResetTimer()
	RunBenchmark(b, func(a int) int {
		return A_log_zerolog(logger, a)
	})
}
