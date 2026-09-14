// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/ass/style"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

const (
	GIANT_FONT_SIZE_WARNING  = 29
	GIANT_FONT_SIZE_CRITICAL = 34
)

func GiantFont() lint.Lint {
	return lint.Lint{
		Name:        "giant-font",
		Description: "fonts that have unusual big size",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			if KaraData.Lyrics[0].ScriptInfo.PlayResX != 0 || KaraData.Lyrics[0].ScriptInfo.PlayResY != 0 {
				return report.Skip("resolution is not 0×0"), nil
			}
			warn := false
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Styles {
				select {
				case <-ctx.Done():
					return report.Abort(), ctx.Err()
				default:
					if after, ok := strings.CutPrefix(line, "Style: "); ok {
						s, err := style.Parse(after)
						if err != nil {
							return report.Abort(), err
						}
						if strings.Contains(strings.ToLower(s.Name), "symbol") {
							// style used by some multi-singer karaoke to make symbols bigger
							continue
						}
						if s.Fontsize >= GIANT_FONT_SIZE_CRITICAL {
							return report.FailCritical("found a style with a big fontsize: " +
								"consider reducing font size " +
								"(it may be hard to identify big text as lyrics to actually sing)"), nil
						}
						if s.Fontsize >= GIANT_FONT_SIZE_WARNING {
							warn = true
						}
					}
				}
			}
			if warn {
				return report.FailWarning("found a style with a big fontsize: " +
					"consider reducing font size " +
					"(it may be hard to identify big text as lyrics to actually sing) "), nil
			}
			return report.Pass(), nil
		},
	}
}
