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
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

func MusicVideoCreditless() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "music-video-creditless",
		Description: "MV with a creditless tag",
		SkipCond: cond.Any{
			cond.NoLyrics{},
			cond.HasNoTagFrom{
				TagType: tag.Songtypes,
				Tags:    []karajson.Tid{songtype.MusicVideo},
				Msg:     "not a music video",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Misc, misc.Creditless) {
				return report.FailCritical("music videos cannot be creditless, remove this tag"), nil
			}

			return report.Pass(), nil
		},
	}
}
