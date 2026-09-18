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

func EolPunctuation() lint.Lint {
	return lint.Lint{
		Name:        "eol-punctuation",
		Description: "non-significant punctuation at end-of-lines",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			hasKaraokeEffectLines := false
			// We need to check on karaoke comment lines, because the karaoke template may create a line per word (karaokes with furigana).
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if (line.Type != lyrics.Format) && (line.Type == lyrics.Comment) && (line.Effect == "karaoke") {
					hasKaraokeEffectLines = true
					l := line.Text.StripTags()
					if strings.HasSuffix(l, ".") || strings.HasSuffix(l, ",") {
						if strings.HasSuffix(l, "...") {
							continue
						}
						return report.FailCritical("remove useless punctuation (`.` or `,`) at end of line"), nil
					}
				}
			}
			if !hasKaraokeEffectLines { // fallback, for old karaokes
				for _, line := range KaraData.Lyrics[0].Events {
					if err := ctx.Err(); err != nil {
						return report.Abort(), err
					}
					if (line.Type != lyrics.Format) && (line.Type != lyrics.Comment) {
						hasKaraokeEffectLines = true
						l := line.Text.StripTags()
						if strings.HasSuffix(l, ".") || strings.HasSuffix(l, ",") {
							if strings.HasSuffix(l, "...") {
								continue
							}
							return report.FailCritical("remove useless punctuation (`.` or `,`) at end of line"), nil
						}
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
