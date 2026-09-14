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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/songtype"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/version"
)

func FullAudioOnlySongtype() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "full-audio-only-songtype",
		Description: "audio only song are considered OP/ED/IN only when they have the same size than actual OP/ED/IN",
		SkipCond: cond.Any{
			cond.HasNoTagFrom{
				TagType: tag.Versions,
				Tags:    []karajson.Tid{version.Full},
				Msg:     "is not a full version",
			},
			cond.HasNoTagFrom{
				TagType: tag.Songtypes,
				Tags:    []karajson.Tid{songtype.AUDIO},
				Msg:     "is not an audio only",
			},
			cond.HasNoTagFrom{
				TagType: tag.Songtypes,
				Tags:    []karajson.Tid{songtype.OP, songtype.ED, songtype.IN},
				Msg:     "is not an OP/ED/IN",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			return report.FailCritical("remove the OP/ED/IN tag"), nil
		}}
}
