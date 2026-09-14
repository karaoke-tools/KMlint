// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
)

func NoLyrics() lint.Lint {
	return lint.Lint{
		Name:        "no-lyrics",
		Description: "missing lyrics file",
		SkipCond:    cond.HasLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if res := KaraData.KaraJson.HasAnyTagFrom(tag.Langs, []karajson.Tid{language.ZXX}); !res {
				return report.FailCritical("no lyrics file, but the media is supposed to have has linguistic content"), nil
			}
			return report.Pass(), nil // no linguistical content
		},
	}
}
