// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/ass/lyrics"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func UnicodeWeirdSpaces() lint.Lint {
	return lint.Lint{
		Name:        "unicode-weird-spaces",
		Description: "detect lyrics file with weird unicode spaces",
		SkipCond: cond.Any{
			cond.NoLyrics{},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if (line.Type != lyrics.Format) && !((line.Type == lyrics.Comment) && (line.Effect != "karaoke")) {
					l := line.Text.StripTags()
					if strings.ContainsRune(l, '\u2005') { // FOUR-PER-EM SPACE
						return report.FailWarning("Found `Four-Per-Em Space` (U+2005): replace it with regular space " +
								"(they may not render correcly on all systems)",
							),
							nil
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
