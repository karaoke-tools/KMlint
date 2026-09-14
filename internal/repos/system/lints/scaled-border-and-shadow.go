// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func ScaledBorderAndShadow() lint.Lint {
	return lint.Lint{
		Name:        "scaled-border-and-shadow",
		Description: "scaled border and shadow not enabled",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			if KaraData.Lyrics[0].ScriptInfo.ScaledBorderAndShadow {
				return report.Pass(), nil
			}
			return report.FailCritical("check the \"Scale border and shadow\" box"), nil
		},
	}
}
