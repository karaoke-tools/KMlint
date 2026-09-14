// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson/tag"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip/cond"
)

func TooManyLangTags() lint.Lint {
	return lint.Lint{
		Name:        "too-many-lang-tags",
		Description: "if more than 2 langs tags, replace them with multilingual tag",
		SkipCond: cond.HasLessTagsThan{
			TagType: tag.Langs,
			Number:  3,
			Msg:     "has not more than 2 lang tags",
		},
		RunFunc: func(ctx context.Context, karaData karadata.KaraData) (report.Report, error) {
			return report.FailCritical("replace lang tags with \"multilingual\""), nil
		},
	}
}
