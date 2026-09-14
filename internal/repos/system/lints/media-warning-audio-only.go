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
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/songtype"
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/warning"
)

// warnings that are related to the media
var mediaWarnings []karajson.Tid = []karajson.Tid{
	warning.R18Media,
	warning.Spoiler,
	warning.Epilepsy,
}

func MediaWarningAudioOnly() lint.Lint {
	return lint.Lint{
		Name:        "media-warning-audio-only",
		Description: "media warning but this is an audio only kara",
		SkipCond: cond.HasNoTagFrom{
			TagType: tag.Songtypes,
			Tags:    []karajson.Tid{songtype.AudioOnly},
			Msg:     "not an audio only",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			for _, w := range mediaWarnings {
				if slices.Contains(KaraData.KaraJson.Data.Tags.Warnings, w) {
					return report.FailCritical("check warning tags " +
						"(maybe a R18-media should be changed to R18-lyrics, maybe a tag should be removed)"), nil
				}
			}
			return report.Pass(), nil
		},
	}
}
