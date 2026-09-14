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
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func SingleCollection() lint.Lint {
	return lint.Lint{
		Name:        "single-collection",
		Description: "multiple collections set",
		SkipCond:    cond.Never{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if len(KaraData.KaraJson.Data.Tags.Collections) != 1 {
				return report.Fail(severity.Critical, "choose the right collection according to rules"), nil
			}
			return report.Pass(), nil
		},
	}
}
