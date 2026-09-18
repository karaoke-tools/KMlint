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

// AutomationAppliedFurugina is a generic version of "`automation-applied` lint" where we only check if at least one
// line has been generated from automation script and no line with "karaoke" effect is uncommented.
func AutomationAppliedFurigana() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "automation-applied-furigana",
		Description: "automation script not applied (song with furigana) ",
		SkipCond: cond.Any{
			cond.NoLyrics{},
			// we skip this lint when this is not a song in kana
			cond.HasNoTagFrom{
				TagType: tag.Collections,
				Tags:    []karajson.Tid{collection.NonLatin},
				Msg:     "song in latin script",
			},
			cond.HasNoTagFrom{
				TagType: tag.Langs,
				Tags:    []karajson.Tid{language.JPN},
				Msg:     "not a japanese song",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			fx := false
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if line.Type == lyrics.Dialogue {
					switch line.Effect {
					case "fx":
						fx = true
					case "karaoke":
						return report.FailCritical("automation script has not been applied"), nil
					}
				}
			}
			if fx {
				return report.Pass(), nil
			}
			return report.FailCritical("automation script has not been applied"), nil
		},
	}
}
