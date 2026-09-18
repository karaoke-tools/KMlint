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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/songtype"
)

func SongtypeConflict() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "songtype-conflict",
		Description: "detects incompatible songtypes",
		SkipCond: cond.HasLessTagsThan{
			TagType: tag.Songtypes,
			Number:  2,
			Msg:     "has a single songtype",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.OT}) {
				return report.FailCritical("songtype \"OT\" is forbidden"), nil
			}
			if KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.AUDIO}) &&
				KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.MV, songtype.AMV}) {
				return report.FailCritical("MV/AMV cannot be audio only"), nil
			}

			counter := 0
			for _, tag := range KaraData.KaraJson.Data.Tags.Songtypes {
				if err := ctx.Err(); err != nil {
					return report.Abort(), err
				}

				// maybe in the future we will move AUDIO into "families",
				// and forbid songs to be both a CS and something else (only allow audio only songs to be CS because"we are not an encyclopedia")
				// then we may simply force this field length to be equal to 1 (maybe directly in KM, by a rule in repo's manifest)
				if !slices.Contains([]karajson.Tid{songtype.AUDIO, songtype.CS}, tag) {
					counter++
					if counter > 1 {
						return report.FailCritical("incompatible songtypes"), nil
					}
				}
			}
			return report.Pass(), nil
		},
	}
}
