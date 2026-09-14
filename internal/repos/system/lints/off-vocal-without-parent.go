// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/version"
)

func OffVocalWithoutParent() lint.Lint {
	return lint.Lint{
		Name:        "off-vocal-without-parent",
		Description: "off vocal but no parent",
		SkipCond: cond.HasNoTagFrom{
			TagType: tag.Versions,
			Tags:    []karajson.Tid{version.OffVocal},
			Msg:     "not an off vocal",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if len(KaraData.KaraJson.Data.Parents) == 0 {
				return report.Fail(severity.Critical, "add the right parent"), nil
			}

			return report.Pass(), nil
		},
	}
}
