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

func FromDisplayType() lint.Lint {
	return lint.Lint{
		Name:        "from-display-type",
		Description: "weird values for from-display-type",
		SkipCond:    cond.Never{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			switch KaraData.KaraJson.Data.FromDisplayType {
			case "series":
				return report.FailWarning("from-display-type is manually set to `" +
					KaraData.KaraJson.Data.FromDisplayType +
					"` but this is already the default"), nil
			case "singergroups":
				if len(KaraData.KaraJson.Data.Tags.Series) == 0 {
					return report.FailWarning("from-display-type is manually set to `" +
						KaraData.KaraJson.Data.FromDisplayType +
						"` but this is already the default"), nil
				}
				fallthrough
			case "singers":
				if len(KaraData.KaraJson.Data.Tags.Series) == 0 && len(KaraData.KaraJson.Data.Tags.Singergroups) == 0 {
					return report.FailWarning("from-display-type is manually set to `" +
						KaraData.KaraJson.Data.FromDisplayType +
						"` but this is already the default"), nil
				}
				fallthrough
			case "songwriters":
				fallthrough
			case "franchises":
				fallthrough
			case "creators":
				fallthrough
			case "":
				return report.Pass(), nil
			}
			return report.FailCritical("from-display-type is set with a weird value: `" + KaraData.KaraJson.Data.FromDisplayType + "`"), nil
		},
	}
}
