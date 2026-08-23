// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package printer

import (
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"os"

	"github.com/karaoke-tools/km-probe/internal/lints"
)

type JsonPrinter struct {
	*BasePrinter
}

func NewJsonPrinter() Printer {
	return JsonPrinter{
		BasePrinter: NewBasePrinter(),
	}
}

func (p JsonPrinter) Encode(ctx context.Context, a *lints.Aggregator) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-p.ready:
		defer p.setReady()
		defer p.aggregatorPool.Put(a)
		if err := json.MarshalWrite(os.Stdout, a, jsontext.WithIndent("  ")); err != nil {
			return err
		}
		if _, err := os.Stdout.WriteString("\n"); err != nil {
			return err
		}
	}
	return nil

}
