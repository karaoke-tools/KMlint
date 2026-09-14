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
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/origin"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/songtype"
)

func SongorderNoOpEd() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "songorder-no-op-ed",
		Description: "songorder is not compatible with this songtype",
		SkipCond: cond.Any{
			cond.HasNoSongorder{},
			cond.HasAnyTagFrom{
				TagType: tag.Songtypes,
				Tags:    []karajson.Tid{songtype.OP, songtype.ED},
				Msg:     "songtype is an OP or ED",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if b := KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.IN, songtype.PV}); b {
				return report.FailWarning("songorder in IS/PV may be justified, but is rare"), nil
			}
			// fanwork + MV/OT: yes this is a strange tag mix, but at least the "songorder" is probably a deliberate choice…
			// playlists might be better for that, but this is probably valid
			if b := KaraData.KaraJson.HasAnyTagFrom(tag.Origins, []karajson.Tid{origin.Fanworks}) &&
				KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.MV, songtype.OT}); b {
				return report.Pass(), nil
			}
			// MV + OVA: probably a serie of music videos
			// playlists might be better for that, but this is probably valid
			if b := KaraData.KaraJson.HasAnyTagFrom(tag.Origins, []karajson.Tid{origin.OVA}) &&
				KaraData.KaraJson.HasAnyTagFrom(tag.Songtypes, []karajson.Tid{songtype.MV}); b {
				return report.Pass(), nil
			}
			return report.FailCritical("remove songorder, or add missing songtype"), nil
		},
	}
}
