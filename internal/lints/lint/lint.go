// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lint

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/skip"
)

type Lint struct {
	Pkg         string
	Name        string
	Description string
	SkipCond    skip.Condition
	RunFunc     func(ctx context.Context, KaraData karadata.KaraData) (report.Report, error)
}

func (l Lint) String() string {
	pkg := "system"
	if l.Pkg != "" {
		pkg = l.Pkg
	}
	name := "undefined"
	if l.Name != "" {
		name = l.Name
	}
	return pkg + "." + name
}

func (l Lint) Run(ctx context.Context, KaraData karadata.KaraData) (report.Report, error) {
	if l.RunFunc == nil {
		return report.Abort(), nil
	}

	if l.SkipCond != nil {
		if skip, msg, err := l.SkipCond.Result(ctx, KaraData); err != nil {
			return report.Abort(), err
		} else if skip {
			return report.Skip(msg), nil
		}
	}

	return l.RunFunc(ctx, KaraData)
}

func (l Lint) Enabled() bool {
	// TODO
	return true
}
