// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package karajson

import "uuid"

// Kid represents a Karaoke ID (Song ID).
type Kid struct {
	uuid.UUID
}

// Tid represents a Tag ID.
type Tid struct {
	uuid.UUID
}

// MustParseTid returns the Tid represented by s.
// It panics if s is not a valid string representation of a KID.
func MustParseTid(s string) Tid {
	return Tid{uuid.MustParse(s)}
}
