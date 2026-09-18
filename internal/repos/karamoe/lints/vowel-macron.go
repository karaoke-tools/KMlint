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

func VowelMacron() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "vowel-macron",
		Description: "ā, ē, ō, ī, ū in lyrics file (JPN romaji only)",
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
					return report.Abort(), ctx.Err()
				}
				if (line.Type != lyrics.Format) && (!strings.HasPrefix(line.Effect, "template")) {
					if strings.ContainsAny(line.Text.StripTags(), "āīūēō") {
						return report.FailWarning("in full japanese song, vowels should not have macron: " +
							"use the appropriate expansion from: aa/ii/uu/ee/ou/oo " +
							"(if this is on a chinese word, make sure to put it in fullcaps)"), nil
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
