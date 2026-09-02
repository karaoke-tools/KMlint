// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package severity

//go:generate go tool stringer -type Severity -linecomment

type Severity int

const (
	// No info
	Unknown Severity = iota // unknown

	// There is something to be corrected
	Critical // critical

	// Maintainer should have a look at it, but may decide to ignore it
	Warning // warning

	// Maintainer can safely ignore this info
	Info // info
)
