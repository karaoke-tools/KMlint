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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/origin"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/songtype"
)

type Creditless struct {
	baselint.BaseLint
	lint.WithDefault
}

var possibleCreditlessOrigin = []karajson.Tid{
	origin.Movie,
	origin.OriginalNetworkAnimation,
	origin.OriginalVideoAnimation,
	origin.TVSpecial,
	origin.TVSeries,
}

func NewCreditless() lint.Lint {
	return &Creditless{
		baselint.New(
			"creditless",
			"can a creditless version be found?",
			cond.Any{
				cond.HasAnyTagFrom{
					TagType: tag.Misc,
					Tags:    []karajson.Tid{misc.Creditless},
					Msg:     "is a creditless version",
				},
				cond.HasAnyTagFrom{
					TagType: tag.Songtypes,
					Tags:    []karajson.Tid{songtype.AUDIO},
					Msg:     "has audio only tag",
				},
				cond.HasNoTagFrom{
					TagType: tag.Songtypes,
					Tags:    []karajson.Tid{songtype.OP, songtype.ED},
					Msg:     "songtype is not OP/ED",
				},
				cond.HasNoTagFrom{
					TagType: tag.Origins,
					Tags:    possibleCreditlessOrigin,
					Msg:     "origin not compatible with creditless",
				},
			},
		),
		baselint.EnabledByDefault{},
	}
}

func (p Creditless) Run(ctx context.Context, KaraData *karadata.KaraData) (report.Report, error) {
	// if the media is already creditless, add the `Creditless`;
	// if a creditless version exists
	// (and is relevant!! see <https://kara.moe/playlist/quand-le-staff-fait-parti-du-generique> for counter-examples),
	// update the media and add the tag
	return report.Fail(severity.Info, "not tagged as creditless"), nil
}
