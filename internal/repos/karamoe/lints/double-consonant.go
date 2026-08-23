// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"strings"

	"github.com/karaoke-tools/km-probe/internal/ass/lyrics"
	"github.com/karaoke-tools/km-probe/internal/karadata"
	"github.com/karaoke-tools/km-probe/internal/karajson"
	"github.com/karaoke-tools/km-probe/internal/karajson/tag"
	"github.com/karaoke-tools/km-probe/internal/lints/lint"
	"github.com/karaoke-tools/km-probe/internal/lints/report"
	"github.com/karaoke-tools/km-probe/internal/lints/report/severity"
	"github.com/karaoke-tools/km-probe/internal/lints/skip/cond"
	"github.com/karaoke-tools/km-probe/internal/repos/karamoe/lints/baselint"
	"github.com/karaoke-tools/km-probe/internal/repos/karamoe/tags/collection"
	"github.com/karaoke-tools/km-probe/internal/repos/system/tags/language"
)

type DoubleConsonant struct {
	baselint.BaseLint
	lint.WithDefault
}

func NewDoubleConsonant() lint.Lint {
	return &DoubleConsonant{
		baselint.New(
			"double-consonant",
			"double consonant in same k-tag (JPN romaji only)",
			cond.Any{
				cond.NoLyrics{},
				cond.HasAnyTagFrom{
					TagType: tag.Collections,
					Tags:    []karajson.Tid{collection.NonLatin},
					Msg:     "non-latin script song",
				},
				cond.HasTagsNotFrom{
					TagType: tag.Langs,
					Tags:    []karajson.Tid{language.JPN},
					Msg:     "not a japanese only version",
				},
			},
		),
		baselint.EnabledByDefault{},
	}
}

var doubleConsonants = []string{
	"kk",
	"gg",
	"ss",
	"zz",
	// avoid matching "mittsu" (みっつ) when pronounced "mi-its’"
	"tta",
	"tte",
	"tti",
	"tto",
	"ttu",
	"dd",
	"nn",
	"bb",
	"pp",
	"mm",
	"rr",
	"cc",
}

func (p DoubleConsonant) Run(ctx context.Context, KaraData *karadata.KaraData) (report.Report, error) {
	// TODO: update this when multi-track drifting is released
	for _, line := range KaraData.Lyrics[0].Events {
		select {
		case <-ctx.Done():
			return report.Abort(), ctx.Err()
		default:
			if (line.Type != lyrics.Format) && (!(line.Type == lyrics.Comment && strings.HasPrefix(line.Effect, "template"))) {
				save := " " // must end with a space to consider the first word of the line as a new word
				for _, syll := range line.Text.TagsSplit {
					select {
					case <-ctx.Done():
						return report.Abort(), ctx.Err()
					default:
						if !strings.HasPrefix(syll, "{") {
							if !strings.HasSuffix(save, " ") { // this is not a new word
								for _, double := range doubleConsonants {
									if strings.HasPrefix(syll, double) {
										return report.Fail(severity.Critical,
											"check for double consonants: "+
												"there is at least an uncorrectly splitted `"+
												strings.TrimSpace(syll)+"`"), nil
									}
								}

							}
							save = syll
						}
					}

				}
			}
		}
	}
	return report.Pass(), nil
}
