// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package version

import "github.com/karaoke-tools/kmlint/internal/karajson"

var (
	// versions
	Cover    = karajson.MustParseTid("03e1e1d2-8641-47b7-bbcb-39a3df9ff21c")
	Full     = karajson.MustParseTid("c2143a7f-6970-450e-8a79-0302db9220a9")
	OffVocal = karajson.MustParseTid("c0cc87b9-55b9-40f0-878a-fbb9e34c151e")

	// versions alias
	Instrumental = OffVocal
)
