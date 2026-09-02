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
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/lints/baselint"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

type MusicVideoCreditless struct {
	baselint.BaseLint
	lint.WithDefault
}

func NewMusicVideoCreditless() lint.Lint {
	return &MusicVideoCreditless{
		baselint.New(
			"music-video-creditless",
			"MV with a creditless tag",
			cond.Any{
				cond.NoLyrics{},
				cond.HasNoTagFrom{
					TagType: tag.Songtypes,
					Tags:    []karajson.Tid{songtype.MusicVideo},
					Msg:     "not a music video",
				},
			},
		),
		baselint.EnabledByDefault{},
	}
}

func (p MusicVideoCreditless) Run(ctx context.Context, KaraData *karadata.KaraData) (report.Report, error) {
	if slices.Contains(KaraData.KaraJson.Data.Tags.Misc, misc.Creditless) {
		return report.Fail(severity.Critical, "music videos cannot be creditless, remove this tag"), nil
	}

	return report.Pass(), nil
}
