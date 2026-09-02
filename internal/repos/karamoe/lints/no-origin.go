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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/lints/baselint"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/songtype"
)

type NoOrigin struct {
	baselint.BaseLint
	lint.WithDefault
}

func NewNoOrigin() lint.Lint {
	return &NoOrigin{
		baselint.New(
			"no-origin",
			"songtype is OP/ED/IN but origin tag is missing",
			cond.Any{
				cond.HasMoreTagsThan{
					TagType: tag.Origins,
					Number:  0,
					Msg:     "has origin",
				},
				cond.HasNoTagFrom{
					TagType: tag.Songtypes,
					Tags:    []karajson.Tid{songtype.OP, songtype.ED, songtype.IN},
					Msg:     "songtype is not OP/ED/IN",
				},
			},
		),
		baselint.EnabledByDefault{},
	}
}

func (p NoOrigin) Run(ctx context.Context, KaraData *karadata.KaraData) (report.Report, error) {
	return report.Fail(severity.Critical, "add the missing origin tag"), nil
}
