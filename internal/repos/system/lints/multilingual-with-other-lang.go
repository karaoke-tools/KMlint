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
	"github.com/karaoke-tools/kmlint/internal/repos/system/tags/language"
)

func MultilingualWithOtherLang() lint.Lint {
	return lint.Lint{
		Name:        "multilingual-with-other-lang",
		Description: "if multilingual tag is applied, no other lang tag should be present",
		SkipCond: cond.Any{
			cond.HasLessTagsThan{
				TagType: tag.Langs,
				Number:  2,
				Msg:     "single lang tag",
			},
			cond.HasNoTagFrom{
				TagType: tag.Langs,
				Tags:    []karajson.Tid{language.MUL},
				Msg:     "has not multilingual tag",
			},
		},
		RunFunc: func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
			return report.FailCritical("check languages tags"), nil

		},
	}
}
