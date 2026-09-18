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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/collection"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
)

// In japanese, `つ` should be timed as a single syllable: `tsu`.
// For example, `ひとつ`(`hitotsu`) should be timed as `hi|to|tsu` and not as `hi|tot|su`.
func WrongTsuSeparation() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "wrong-tsu-separation",
		Description: "`t|su` separation is not correct (JPN romaji only)",
		SkipCond: cond.Any{
			cond.NoLyrics{},
			cond.HasAnyTagFrom{
				TagType: tag.Collections,
				Tags:    []karajson.Tid{collection.NonLatin},
				Msg:     "non-latin script song",
			},
			cond.HasTagsNotFrom{
				TagType: tag.Langs,
				Tags:    []karajson.Tid{language.JPN},
				Msg:     "not a japanese only song",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if (line.Type != lyrics.Format) && (!strings.HasPrefix(line.Effect, "template")) {
					ok := false
					for _, syll := range line.Text.TagsSplit {
						if err := ctx.Err(); err != nil {
							return report.Abort(), err
						}
						if !strings.HasPrefix(syll, "{") {
							if strings.HasSuffix(syll, "t") {
								ok = true
							} else if ok && strings.HasPrefix(syll, "su") {
								return report.FailCritical("`tsu` must be timed as a single syllable"), nil
							} else {
								ok = false
							}
						}
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
