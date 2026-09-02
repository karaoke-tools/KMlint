// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package status

type Status int

//go:generate go tool stringer -type Status -linecomment
const (
	Unknown   Status = iota // unknown status
	Completed               // completed
	Aborted                 // aborted
	Skipped                 //skipped
)
