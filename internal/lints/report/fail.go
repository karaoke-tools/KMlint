// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package report

import (
	"github.com/karaoke-tools/kmlint/internal/lints/report/result"
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/report/status"
)

// When the issue is detected
func Fail(severity severity.Severity, message string) *report {
	r := reportPool.Get().(*report)
	r.severity = severity
	r.message = message // indicate what action must be done
	r.status = status.Completed
	r.result = result.Failed
	return r
}

// Fails with [severity.Critical]
func FailCritical(message string) *report {
	return Fail(severity.Critical, message)
}

// Fails with [severity.Warning]
func FailWarning(message string) *report {
	return Fail(severity.Warning, message)
}

// Fails with [severity.Info]
func FailInfo(message string) *report {
	return Fail(severity.Info, message)
}
