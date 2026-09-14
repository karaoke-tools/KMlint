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

func Resolution() lint.Lint {
	return lint.Lint{
		Name:        "resolution",
		Description: "resolution not set to 0×0",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			if KaraData.Lyrics[0].ScriptInfo.PlayResX == 0 && KaraData.Lyrics[0].ScriptInfo.PlayResY == 0 {
				return report.Pass(), nil
			}
			return report.FailCritical("update resolution to be 0×0 (and check style size)"), nil
		},
	}
}
