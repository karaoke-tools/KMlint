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
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
)

func SpaceBeforeDoublePunctuation() lint.Lint {
	return lint.Lint{
		Name:        "space-before-double-punctuation",
		Description: "space before double punctuation (JPN/ENG only)",
		SkipCond: cond.Any{
			cond.NoLyrics{},
			cond.HasTagsNotFrom{
				TagType: tag.Langs,
				Tags:    []karajson.Tid{language.JPN, language.ENG},
				Msg:     "non english/japanese language",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if (line.Type != lyrics.Format) && !((line.Type == lyrics.Comment) && (line.Effect != "karaoke")) {
					l := line.Text.StripTags()
					if strings.Contains(l, " ?") || strings.Contains(l, " !") ||
						strings.Contains(l, " ?") || strings.Contains(l, " !") { // non-breakable space
						return report.FailCritical("remove space before `?`/`!`"), nil
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
