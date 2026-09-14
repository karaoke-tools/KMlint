// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"slices"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/warning"
)

func LyricsWarningZXX() lint.Lint {
	return lint.Lint{
		Name:        "lyrics-warning-zxx",
		Description: "lyrics warning, but there is no linguistical content",
		SkipCond: cond.HasNoTagFrom{
			TagType: tag.Warnings,
			Tags:    []karajson.Tid{warning.R18Lyrics},
			Msg:     "no lyrics-warning tag",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Langs, language.ZXX) {
				return report.FailCritical("check if lyrics warning is relevant, and if the Langs field is set"), nil
			}
			return report.Pass(), nil
		},
	}
}
