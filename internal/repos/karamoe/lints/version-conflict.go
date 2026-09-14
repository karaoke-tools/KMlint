// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"slices"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/version"
)

func VersionConflict() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "version-conflict",
		Description: "incompatible version tags",
		SkipCond: cond.HasLessTagsThan{
			TagType: tag.Versions,
			Number:  2,
			Msg:     "has not multiple version tags",
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Versions, version.Short) &&
				slices.Contains(KaraData.KaraJson.Data.Tags.Versions, version.Full) {
				return report.FailCritical("song cannot be both a short version and a full version at the same time"), nil
			}
			if slices.Contains(KaraData.KaraJson.Data.Tags.Versions, version.Acoustic) &&
				slices.Contains(KaraData.KaraJson.Data.Tags.Versions, version.Metal) {
				return report.FailCritical("song cannot be both a metal version and an acoustic version at the same time"), nil
			}
			return report.Pass(), nil
		},
	}
}
