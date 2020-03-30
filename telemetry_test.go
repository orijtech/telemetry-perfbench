// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"golang.org/x/net/context"
	"golang.org/x/tools/telemetry/event"
	"golang.org/x/tools/telemetry/export"
)

var (
	aKey    = event.NewIntKey("a", "")
	bKey    = event.NewStringKey("b", "")
	lenBKey = event.NewIntKey("len(b)", "")
)

func a_log_telemetry(ctx context.Context, a int) int {
	if a > 0 {
		event.Print1(ctx, "a > 0", aKey.Of(a))
		_ = 10 * 12
	}
	event.Print(ctx, "calling b")
	return b_log_telemetry(ctx, "Called from A")
}

func b_log_telemetry(ctx context.Context, b string) int {
	b = strings.ToUpper(b)
	event.Print1(ctx, "b uppercased, so lowercased", lenBKey.Of(len(b)))
	if len(b) > 1024 {
		b = strings.ToLower(b)
		event.Print1(ctx, "b > 1024, so lowercased", bKey.Of(b))
	}
	return len(b)
}

func BenchmarkLoggingTelemetryOff(b *testing.B) {
	event.SetExporter(nil)
	ctx := context.Background()
	RunBenchmark(b, func(a int) int {
		return a_log_telemetry(ctx, a)
	})
}

func BenchmarkLoggingTelemetryIgnore(b *testing.B) {
	event.SetExporter(func(ctx context.Context, ev event.Event, tagMap event.TagMap) context.Context {
		return ctx
	})
	ctx := context.Background()
	RunBenchmark(b, func(a int) int {
		return a_log_telemetry(ctx, a)
	})
}

func BenchmarkLoggingTelemetryDiscard(b *testing.B) {
	event.SetExporter(export.LogWriter(ioutil.Discard, false))
	ctx := context.Background()
	RunBenchmark(b, func(a int) int {
		return a_log_telemetry(ctx, a)
	})
}
