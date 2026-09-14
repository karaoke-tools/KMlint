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
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

func AudioOnlyWithFamilies() lint.Lint {
	return lint.Lint{
		Name:        "audio-only-with-families",
		Description: "media content tag including both audio only tag and other tags at the same time",
		SkipCond: cond.HasNoTagFrom{
			TagType: tag.Songtypes,
			Tags:    []karajson.Tid{songtype.AudioOnly},
			Msg:     "not an audio only",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if len(KaraData.KaraJson.Data.Tags.Families) > 0 {
				return report.FailCritical("an audio only media cannot have a content type (family)"), nil
			}
			return report.Pass(), nil
		},
	}
}
