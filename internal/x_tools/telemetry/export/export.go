// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package export is a stub to expose parts of
// golang.org/x/tools/internal/telemetry/export to the
// benchmark.
package export

import (
	iex "golang.org/x/tools/internal/telemetry/export"
)

var (
	LogWriter = iex.LogWriter
)
