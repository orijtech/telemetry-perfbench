package log_test

import (
	"io/ioutil"
	stdlog "log"
	"strings"
	"testing"
)

func init() {
	stdlog.SetOutput(ioutil.Discard)
}

func A_log_stdlib(a int) int {
	if a > 0 {
		stdlog.Printf("a > 0 where a=%d", a)
		_ = 10 * 12
	}
	stdlog.Print("calling b")
	return B_log_stdlib("Called from A")
}

func B_log_stdlib(b string) int {
	b = strings.ToUpper(b)
	stdlog.Printf("b uppercased, so lowercased where len_b=%d", len(b))
	if len(b) > 1024 {
		b = strings.ToLower(b)
		stdlog.Printf("b > 1024, so lowercased where b=%s", b)
	}
	return len(b)
}
func BenchmarkLoggingStdlib(b *testing.B) {
	RunBenchmark(b, A_log_stdlib)
}
