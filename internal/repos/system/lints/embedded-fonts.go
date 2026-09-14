// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func EmbeddedFonts() lint.Lint {
	return lint.Lint{
		Name:        "embedded-fonts",
		Description: "lyrics file embeds fonts",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			if KaraData.Lyrics[0].Fonts {
				return report.FailCritical(
					"lyrics file embeds fonts; consider using standard fonts instead because fonts embedding creates big lyrics file"), nil
			}
			return report.Pass(), nil
		},
	}
}
