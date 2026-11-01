// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/ass/lyrics"
	"github.com/karaoke-tools/kmlint/internal/ass/style"
	"github.com/karaoke-tools/kmlint/internal/ass/style/colour"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func StyleBadContrast() lint.Lint {
	return lint.Lint{
		Name:        "style-bad-contrast",
		Description: "detects some known bad contrast styles",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			badStyles := make([]string, 0, len(KaraData.Lyrics[0].Styles)-1)
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
					if s.PrimaryColour == colour.Red || s.SecondaryColour == colour.Red {
						badStyles = append(badStyles, s.Name)
					}
				}
			}
			if len(badStyles) == 0 {
				return report.Pass(), nil
			}
			// TODO: update this when multi-track drifting is released
			badStylesCurrent := make([]string, 0, len(badStyles))

			// list of used styles
			for _, line := range KaraData.Lyrics[0].Events {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}
				if line.Type != lyrics.Dialogue {
					continue
				}
				for _, style := range line.Styles() {
					if err := ctx.Err(); err != nil {
						return report.Abort(), err
					}
					if slices.Contains(badStylesCurrent, style) {
						continue
					}
					if slices.Contains(badStyles, style) {
						badStylesCurrent = append(badStylesCurrent, style)
					}
				}
			}
			if len(badStylesCurrent) > 0 {
				return report.FailCritical(fmt.Sprintf("bad styles in-use: %s", badStylesCurrent)), nil
			}
			return report.Pass(), nil
		},
	}
}
