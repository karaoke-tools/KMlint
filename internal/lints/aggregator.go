// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package lints

import (
	"context"
	"time"

	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"
	"github.com/karaoke-tools/kmlint/internal/lints/lint"
	"github.com/karaoke-tools/kmlint/internal/lints/report"
	"github.com/karaoke-tools/kmlint/internal/lints/report/result"
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/report/status"
)

type Aggregator struct {
	// Identification of the song
	Repository string       `json:"repository"`
	Songname   string       `json:"songname"`
	Kid        karajson.Kid `json:"kid"`
	CreatedAt  time.Time    `json:"created-at"`
	ModifiedAt time.Time    `json:"modified-at"`
	Year       int          `json:"year"`
	// [Lints] report direct features of the song based on metadata, lyrics, etc.
	// They can be used to detect common mistakes.
	Lints   []lint.Lint              `json:"-"`
	Reports map[string]report.Report `json:"reports"`
	Stats   *Stats                   `json:"statistics"`
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		Lints:   Enabled(),
		Reports: make(map[string]report.Report),
		Stats:   &Stats{},
	}
}

func (a *Aggregator) Reset(basedir string, karaJson karajson.KaraJson) {
	// recycle reports & analysis memory
	for _, v := range a.Reports {
		v.Delete()
	}
	// empty the map
	clear(a.Reports)
	a.Repository = karaJson.Data.Repository
	a.Songname = karaJson.Data.Songname
	a.Kid = karaJson.Data.Kid
	a.CreatedAt = karaJson.Data.CreatedAt
	a.ModifiedAt = karaJson.Data.ModifiedAt
	a.Year = karaJson.Data.Year
	a.Stats.Reset()
}

type reportWithName struct {
	name string
	r    report.Report
}

func (a *Aggregator) Run(ctx context.Context, KaraData karadata.KaraData) error {
	if err := ctx.Err(); err != nil {
		// if a.Lints is empty, context would not be checked otherwise
		return err
	}
	ch := make(chan reportWithName)
	// start lints
	for _, l := range a.Lints {
		if err := ctx.Err(); err != nil {
			return err
		}
		go func(ctx context.Context, l lint.Lint, ch chan<- reportWithName) error {
			r, err := l.Run(ctx, KaraData)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- reportWithName{name: l.String(), r: r}:
				return err
			}
		}(ctx, l, ch)
	}
	// get result of lints
	for range a.Lints {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case r := <-ch:
			a.Reports[r.name] = r.r
			switch r.r.Status() {
			case status.Completed:
				switch r.r.Result() {
				case result.Passed:
					a.Stats.Passed += 1
				case result.Failed:
					switch r.r.Severity() {
					case severity.Info:
						a.Stats.FailedInfo += 1
					case severity.Warning:
						a.Stats.FailedWarning += 1
					case severity.Critical:
						a.Stats.FailedCritical += 1
					}
				}
			case status.Aborted:
				a.Stats.Aborted += 1
			case status.Skipped:
				a.Stats.Skipped += 1
			}

		}
	}
	return nil
}
