// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/ass/style"
	"github.com/karaoke-tools/kmlint/internal/ass/style/colour"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func StyleBlackBorder() lint.Lint {
	return lint.Lint{
		Name:        "style-black-border",
		Description: "detects non-black border",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Styles {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if strings.HasPrefix(line, "Style: ") && !strings.Contains(line, "-furigana") {
					// we don't care about furigana styles for the now
					s, err := style.Parse(strings.TrimPrefix(line, "Style: "))
					if err != nil {
						return report.Abort(), err
					}
					if s.OutlineColour != colour.Black {
						// border color must be black
						return report.FailWarning("outline must be black " +
							"(this lint can only check if this is pure black, nuances of black might be okay"), nil
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
