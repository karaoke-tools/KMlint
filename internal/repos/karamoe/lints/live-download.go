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
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/collection"
	"github.com/karaoke-tools/kmlint/internal/repos/karamoe/tags/misc"
)

// State of "no live download" collections as of 2025-01-06
var collectionsNoLiveDownload = []karajson.Tid{
	collection.Asia,
	collection.NonLatin,
	collection.West,
}

func isNoLiveDownloadCollection(collection karajson.Tid) bool {
	return slices.Contains(collectionsNoLiveDownload, collection)
}

// LiveDownload only check for hardcoded collections and "unavailable" tag
// because checking each tag may be long when probing the full repository.
func LiveDownload() lint.Lint {
	return lint.Lint{
		Pkg:         PKG_NAME,
		Name:        "live-download",
		Description: "is hardsub available?",
		SkipCond:    cond.Never{},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			if slices.Contains(KaraData.KaraJson.Data.Tags.Misc, misc.Unavailable) {
				return report.FailInfo("not available for live download"), nil
			}
			if slices.ContainsFunc(KaraData.KaraJson.Data.Tags.Collections, isNoLiveDownloadCollection) {
				return report.FailInfo("not available for live download"), nil
			}
			return report.Pass(), nil
		},
	}
}
