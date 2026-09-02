// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package genre

import "github.com/karaoke-tools/kmlint/internal/karajson"

var (
	// genres
	BoysLove         = karajson.MustParseTid("724d2d80-a531-4579-bdaf-4a0db9a313d9")
	GirlsLove        = karajson.MustParseTid("76378421-153e-410d-a62c-2db43d5ed5fd")
	Idol             = karajson.MustParseTid("caa86df2-0d59-474b-885c-f240a9e891b0")
	Isekai           = karajson.MustParseTid("4b80eca6-f44f-4262-a8d0-b6879d6b70cb")
	MagicalGirl      = karajson.MustParseTid("b84e28f9-db1e-447f-b339-7954c3592523")
	Mecha            = karajson.MustParseTid("f60ba57d-5ef4-49e0-b93b-d0dcbcff6592")
	Otome            = karajson.MustParseTid("7025ae3c-7bd6-4787-8912-2c15063343bb")
	Shoujo           = karajson.MustParseTid("8b6ace6f-a59e-4740-b3ff-e1618720383c")
	Shounen          = karajson.MustParseTid("a0aeef4a-6428-45ff-a6e1-468b595930c2")
	Tokusatsu        = karajson.MustParseTid("d3fac9ab-630d-402d-a392-3d2450c3e62e")
	VocalSynthesizer = karajson.MustParseTid("45623cae-6d68-4304-a49b-896d1d6f4580")

	// genres alias
	BL       = BoysLove
	GL       = GirlsLove
	MAGIC    = MagicalGirl
	OTO      = Otome
	SYNTH    = VocalSynthesizer
	TKU      = Tokusatsu
	Vocaloid = VocalSynthesizer
)
