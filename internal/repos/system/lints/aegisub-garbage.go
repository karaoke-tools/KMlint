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

func AegisubGarbage() lint.Lint {
	return lint.Lint{
		Name:        "aegisub-garbage",
		Description: "Aegisub Project Garbage section has not been removed",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			if KaraData.Lyrics[0].AegisubGarbage {
				return report.FailInfo("Aegisub Project Garbage section has not been removed; " +
					"if you are integrating this song make sure to enable the \"cleanup lyrics\" function in Karaoke Mugen "), nil
			}
			return report.Pass(), nil
		},
	}
}
