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
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/version"
)

var versionsWithoutParentInfo = []karajson.Tid{
	version.Cover,
	version.OffVocal,
	version.NonLatin,
}

var versionsWithoutParentCritical = []karajson.Tid{
	version.Acoustic,
	version.Alternative,
	version.Full,
	version.Metal,
	version.Short,
}

func isVersionWithoutParentCritical(versionType karajson.Tid) bool {
	return slices.Contains(versionsWithoutParentCritical, versionType)
}

func isVersionWithoutParentInfo(versionType karajson.Tid) bool {
	return slices.Contains(versionsWithoutParentInfo, versionType)
}

func AltVersionWithoutParent() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "alt-version-without-parent",
		Description: "version tag, but there is no parent song",
		SkipCond: cond.Any{
			cond.HasParent{},
			cond.HasEmptyTagtype{
				TagType: tag.Versions,
				Msg:     "not an alt version",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.ContainsFunc(KaraData.KaraJson.Data.Tags.Versions, isVersionWithoutParentCritical) {
				return report.FailCritical("check if a potential parent exists, or if the version tag is relevant"), nil
			}

			if slices.ContainsFunc(KaraData.KaraJson.Data.Tags.Versions, isVersionWithoutParentInfo) {
				return report.FailInfo("has a version tag but no parent: if the parent exist, don't forget to add it"), nil
			}
			return report.Pass(), nil
		},
	}
}
