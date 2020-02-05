package log_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func a_log_logrus(logger *logrus.Logger, a int) int {
	if a > 0 {
		logger.Infof("a > 0 where a=%d", a)
		_ = 10 * 12
	}
	logger.Info("calling b")
	return b_log_logrus(logger, "Called from A")
}

func b_log_logrus(logger *logrus.Logger, b string) int {
	b = strings.ToUpper(b)
	logger.Infof("b uppercased, so lowercased where len_b=%d", len(b))
	if len(b) > 1024 {
		b = strings.ToLower(b)
		logger.Infof("b > 1024, so lowercased where b=%s", b)
	}
	return len(b)
}

func BenchmarkLoggingLogrus(b *testing.B) {
	logger := logrus.Logger{
		Out:   ioutil.Discard,
		Level: logrus.InfoLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:  true,
			DisableSorting: true,
			DisableColors:  true,
		},
	}
	b.ResetTimer()
	RunBenchmark(b, func(a int) int {
		return a_log_logrus(&logger, a)
	})

}
