// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package result

type Result int

//go:generate go tool stringer -type Result -linecomment
const (
	Unknown Result = iota // unknown result
	Passed                // passed
	Failed                // failed
)
