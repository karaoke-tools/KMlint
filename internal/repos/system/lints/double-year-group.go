// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"slices"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/year"
)

// we cannot just checking the len of the group field,
// because on some repositories it is used for more than only years
var yearsGroup []karajson.Tid = []karajson.Tid{
	year.Y1950,
	year.Y1960,
	year.Y1970,
	year.Y1980,
	year.Y1990,
	year.Y2000,
	year.Y2010,
	year.Y2020,
}

func DoubleYearGroup() lint.Lint {
	return lint.Lint{
		Name:        "double-year-group",
		Description: "double-year-group",
		SkipCond:    cond.Never{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if len(KaraData.KaraJson.Data.Tags.Groups) < 2 {
				return report.Pass(), nil
			}
			ok := false
			for _, group := range KaraData.KaraJson.Data.Tags.Groups {
				if found := slices.Contains(yearsGroup, group); found && ok {
					return report.FailCritical("remove all years group and save to apply (let years hooks do their job)"), nil
				} else if found {
					ok = true
				}
			}
			return report.Pass(), nil
		},
	}
}
