package log_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func A_log_zap_sugared(logger *zap.SugaredLogger, a int) int {
	if a > 0 {
		logger.Infof("a > 0 where a=%d", a)
		_ = 10 * 12
	}
	logger.Info("calling b")
	return B_log_zap_sugared(logger, "Called from A")
}

func B_log_zap_sugared(logger *zap.SugaredLogger, b string) int {
	b = strings.ToUpper(b)
	logger.Infof("b uppercased, so lowercased where len_b=%d", len(b))
	if len(b) > 1024 {
		b = strings.ToLower(b)
		logger.Infof("b > 1024, so lowercased where b=%s", b)
	}
	return len(b)
}

type DummyWriteSyncer struct {
}

func (d *DummyWriteSyncer) Sync() error {
	return nil
}

func (d *DummyWriteSyncer) Write(p []byte) (int, error) {
	return ioutil.Discard.Write(p)
}

func newZapLogger(lvl zapcore.Level) *zap.Logger {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapcore.NanosDurationEncoder
	ec.EncodeTime = zapcore.EpochNanosTimeEncoder
	enc := zapcore.NewJSONEncoder(ec)
	return zap.New(zapcore.NewCore(
		enc,
		&DummyWriteSyncer{},
		lvl,
	))
}

func BenchmarkLoggingZapSugared(b *testing.B) {
	logger := newZapLogger(zap.DebugLevel).Sugar()
	RunBenchmark(b, func(a int) int {
		return A_log_zap_sugared(logger, a)
	})
}
