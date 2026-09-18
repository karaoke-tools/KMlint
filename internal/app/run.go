// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package app

import (
	"context"

	"github.com/karaoke-tools/kmlint/internal/app/printer"
	"github.com/karaoke-tools/kmlint/internal/karadata"
	"github.com/karaoke-tools/kmlint/internal/karajson"

	"github.com/sirupsen/logrus"
)

// Parse the song, run lints, and display result
// p is the filepath to the .kara.json file
func RunOnFile(ctx context.Context, repo Repository, p string, pr printer.Printer) error {
	karaJson, err := karajson.ParseFile(p)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"repository": repo.Name,
			"filepath":   p,
		}).Error("Could not parse karajson file")
		return err
	}
	karaData, err := karadata.New(repo.BaseDir, karaJson)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"repository": repo.Name,
			"filepath":   p,
		}).Error("Could not create karadata")
		return err
	}

	if err := karaData.ParseLyrics(ctx); err != nil {
		if ctx.Err() == nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"repository": repo.Name,
				"filepath":   p,
			}).Error("Could not parse lyrics")
		}
		return err
	}
	aggregator := pr.Aggregator()
	aggregator.Reset(repo.BaseDir, karaJson)
	if err := aggregator.Run(ctx, karaData); err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			logrus.WithError(err).WithFields(logrus.Fields{
				"repository": repo.Name,
				"filepath":   p,
			}).Error("Lint aggregator failure")
			return err
		}
	}
	if err := pr.Encode(ctx, aggregator); err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			logrus.WithError(err).WithFields(logrus.Fields{
				"repository": repo.Name,
				"filepath":   p,
			}).Error("Could not print aggregator result")
			return err
		}
	}
	return nil
}
