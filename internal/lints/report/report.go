// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package report

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"maps"
	"slices"
	"sync"

	"github.com/karaoke-tools/kmlint/internal/lints/report/result"
	"github.com/karaoke-tools/kmlint/internal/lints/report/severity"
	"github.com/karaoke-tools/kmlint/internal/lints/report/status"
)

type Report interface {
	json.MarshalerTo

	Result() result.Result // true: passed, false: failed
	Status() status.Status // completed, aborted, skipped, etc.
	Severity() severity.Severity
	Message() string

	// Delete should be used when the Report is no longer useful.
	// This allows to recycle the memory for future usage.
	Delete()
}

func (r *report) MarshalJSONTo(enc *jsontext.Encoder) error {
	m := map[string]string{
		"result":   r.result.String(),
		"severity": r.severity.String(),
		"status":   r.status.String(),
	}
	if r.message != "" {
		m["message"] = r.message
	}

	if err := enc.WriteToken(jsontext.BeginObject); err != nil {
		return err
	}
	keys := maps.Keys(m)
	deterministic, _ := json.GetOption(enc.Options(), json.Deterministic)
	if deterministic {
		keys = slices.Values(slices.Sorted(keys))
	}

	for k := range keys {
		if err := json.MarshalEncode(enc, k); err != nil {
			return err
		}
		if err := json.MarshalEncode(enc, m[k]); err != nil {
			return err
		}
	}

	if err := enc.WriteToken(jsontext.EndObject); err != nil {
		return err
	}
	return nil
}

type report struct {
	status   status.Status
	result   result.Result
	message  string
	severity severity.Severity

	// when not set, the pool is not used
	recycleAfterUse bool
}

// pool to recycle memory
var reportPool = sync.Pool{
	New: func() any {
		return &report{
			recycleAfterUse: true,
		}
	},
}

// Delete should be used when the report is no longer useful.
// This allows to recycle the memory for future usage.
func (r *report) Delete() {
	if !r.recycleAfterUse {
		return
	}
	r.severity = severity.Unknown
	r.status = status.Unknown
	r.result = result.Unknown
	r.message = ""
	reportPool.Put(r)
}

func (r *report) Status() status.Status {
	return r.status
}

func (r *report) Result() result.Result {
	return r.result
}

func (r *report) Severity() severity.Severity {
	return r.severity
}

func (r *report) Message() string {
	return r.message
}
