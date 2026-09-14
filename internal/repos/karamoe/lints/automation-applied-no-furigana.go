// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/ass/lyrics"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/collection"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
)

func AutomationAppliedNoFurigana() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "automation-applied-no-furigana",
		Description: "automation script not applied (song without furigana)",
		SkipCond: cond.Any{
			cond.NoLyrics{},
			cond.All{
				// we skip this lint when song is in furigana (non-latin with japanese)
				cond.HasAnyTagFrom{
					TagType: tag.Collections,
					Tags:    []karajson.Tid{collection.NonLatin},
					Msg:     "song with latin script",
				},
				cond.HasAnyTagFrom{
					TagType: tag.Langs,
					Tags:    []karajson.Tid{language.JPN},
					Msg:     "japanese song",
				},
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			fx := 0
			karaoke := 0
			for _, line := range KaraData.Lyrics[0].Events {
				select {
				case <-ctx.Done():
					return report.Abort(), ctx.Err()
				default:
					if line.Type == lyrics.Comment && line.Effect == "karaoke" {
						karaoke++
					} else if line.Type == lyrics.Dialogue {
						switch line.Effect {
						case "fx":
							fx++
						case "karaoke":
							return report.FailCritical("automation script has not been applied"), nil
						}
					}
				}
			}
			if fx == 0 || karaoke != fx {
				return report.FailCritical("automation script has not been applied"), nil
			}
			return report.Pass(), nil
		},
	}
}
