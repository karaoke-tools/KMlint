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
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
)

func VideoContainerWithAudioOnlyTag() lint.Lint {
	return lint.Lint{
		Name:        "video-container-with-audio-only-tag",
		Description: "video container, but audio only tag",
		SkipCond:    cond.HasNotVideoExtension{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Songtypes, songtype.AudioOnly) {
				return report.Fail(severity.Critical,
						"if this is a still image replace media with an audio container, "+
							"otherwise replace audio only tag with add appropriate family tag"),
					nil
			}
			return report.Pass(), nil
		},
	}
}
