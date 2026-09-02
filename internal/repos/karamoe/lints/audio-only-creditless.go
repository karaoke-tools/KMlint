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
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

type AudioOnlyCreditless struct {
	baselint.BaseLint
	lint.WithDefault
}

func NewAudioOnlyCreditless() lint.Lint {
	return &AudioOnlyCreditless{
		baselint.New(
			"audio-only-creditless",
			"audio only songs cannot be creditless",
			cond.Any{
				cond.HasNoTagFrom{
					TagType: tag.Songtypes,
					Tags:    []karajson.Tid{songtype.AUDIO},
					Msg:     "song is not audio only",
				},
				cond.HasNoTagFrom{
					TagType: tag.Misc,
					Tags:    []karajson.Tid{misc.Creditless},
					Msg:     "song is not tagged as creditless",
				},
			},
		),
		baselint.EnabledByDefault{},
	}
}

func (p AudioOnlyCreditless) Run(ctx context.Context, karaData *karadata.KaraData) (report.Report, error) {
	return report.Fail(severity.Critical, "remove the creditless tag, or update the media content"), nil
}
