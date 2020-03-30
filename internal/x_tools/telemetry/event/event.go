// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package event is a stub to expose parts of
// golang.org/x/tools/internal/telemetry/event to the
// benchmark.
package event

import (
	iev "golang.org/x/tools/internal/telemetry/event"
)

type (
	Event  = iev.Event
	TagMap = iev.TagMap
)

var (
	Print        = iev.Print
	Print1       = iev.Print1
	NewIntKey    = iev.NewIntKey
	NewStringKey = iev.NewStringKey
	SetExporter  = iev.SetExporter
)
