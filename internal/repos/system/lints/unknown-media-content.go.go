// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"slices"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

func UnknownMediaContent() lint.Lint {
	return lint.Lint{
		Name:        "unknown-media-content",
		Description: "missing content tag",
		SkipCond:    cond.Never{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if len(KaraData.KaraJson.Data.Tags.Families) > 0 {
				return report.Pass(), nil
			}
			if slices.Contains(KaraData.KaraJson.Data.Tags.Songtypes, songtype.AudioOnly) {
				return report.Pass(), nil
			}
			return report.FailCritical("indicate the media content type (animation, real, audio only)"), nil
		},
	}
}
