// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package version

import "github.com/karaoke-tools/kmlint/internal/karajson"

var (
	// versions
	Acoustic    = karajson.MustParseTid("dd5e063e-8a40-478e-a4c5-d477ce8f13f7")
	Alternative = karajson.MustParseTid("9f63c359-d1bf-425a-bcf5-b245c2c9211d")
	Cover       = karajson.MustParseTid("03e1e1d2-8641-47b7-bbcb-39a3df9ff21c")
	Full        = karajson.MustParseTid("c2143a7f-6970-450e-8a79-0302db9220a9")
	Metal       = karajson.MustParseTid("188a5c46-63ff-4e9f-89e4-763468b6ea4a")
	NonLatin    = karajson.MustParseTid("12d21e9c-3427-4677-a5bd-1301a0f7358a")
	OffVocal    = karajson.MustParseTid("c0cc87b9-55b9-40f0-878a-fbb9e34c151e")
	Short       = karajson.MustParseTid("d125c80a-86c1-4321-88e9-07b7800679b6")

	// versions alias
	Instrumental = OffVocal
)
