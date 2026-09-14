// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"slices"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
)

func LongTagOnShortMedia() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "long-tag-on-short-media",
		Description: "long tag added manually",
		SkipCond:    cond.GreaterMediaDuration{Duration: 300},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Misc, misc.Long) {
				return report.FailCritical("remove long tag"), nil
			}
			return report.Pass(), nil
		},
	}
}
