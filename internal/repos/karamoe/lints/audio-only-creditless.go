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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

func AudioOnlyCreditless() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "audio-only-creditless",
		Description: "audio only songs cannot be creditless",
		SkipCond: cond.Any{
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
		RunFunc: func(ctx context.Context, karaData karadata.KaraData) (report.Report, error) {
			return report.FailCritical("remove the creditless tag, or update the media content"), nil
		},
	}
}
