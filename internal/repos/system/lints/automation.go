// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/ass/lyrics"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func Automation() lint.Lint {
	return lint.Lint{
		Name:        "automation",
		Description: "missing automation script",
		SkipCond:    cond.NoLyrics{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			// TODO: update this when multi-track drifting is released
			for _, line := range KaraData.Lyrics[0].Events {
				select {
				case <-ctx.Done():
					return report.Abort(), ctx.Err()
				default:
					if line.Type == lyrics.Comment {
						return report.Pass(), nil
					}
				}
			}
			return report.FailCritical("missing automation line in the lyrics file"), nil
		},
	}
}
