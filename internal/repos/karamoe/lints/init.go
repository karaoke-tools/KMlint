// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"github.com/karaoke-tools/kmlint/internal/lints"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
)

const PKG_NAME = "kara-moe"

func init() {
	lints.Register([]lint.Lint{
		AltVersionWithoutParent(),
		AudioOnlyCreditless(),
		AutomationAppliedFurigana(),
		AutomationAppliedNoFurigana(),
		Creditless(),
		DoubleConsonant(),
		FullAudioOnlyOrigin(),
		FullAudioOnlySongtype(),
		LiveDownload(),
		LongTagOnShortMedia(),
		MusicVideoCreditless(),
		NoOrigin(),
		SongorderNoOpEd(),
		SongtypeConflict(),
		StyleSingleWhite(),
		VersionConflict(),
		VowelMacron(),
		WrongTsuSeparation(),
	})
}
