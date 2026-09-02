// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package info

import (
	"context"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"
	"os"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/app/ansi"
	"github.com/karaoke-tools/kmlint/internal/app/setup"
	"github.com/karaoke-tools/kmlint/internal/lints"

	"github.com/urfave/cli/v3"
)

type ListLints struct {
	setup.Setup
}

func FromCommand(command *cli.Command) ListLints {
	return ListLints{Setup: setup.FromCommand(command)}
}

type prb struct {
	Name          string `json:"name"`
	Desc          string `json:"description"`
	Enabled       bool   `json:"enabled"`
	EnabledString string `json:"-"`
}

func enabledString(enabled bool) string {
	if enabled {
		return "enabled"
	}
	return "disabled"
}

type color int

const (
	green color = iota
	red
	noColor
)

func (p prb) Write(out io.Writer, namelen int, enabledlen int, b *strings.Builder, underline bool, color color) error {
	defer b.Reset()
	switch color {
	case green:
		b.WriteString(ansi.Green)
	case red:
		b.WriteString(ansi.Red)
	}
	// Enabled
	if underline {
		b.WriteString(ansi.Underline)
	}
	b.WriteString(p.EnabledString)
	if underline {
		b.WriteString(ansi.Reset)
	}
	for range enabledlen - len(p.EnabledString) {
		b.WriteString(" ")
	}

	b.WriteString("\t")

	// Name
	if underline {
		b.WriteString(ansi.Underline)
	}
	b.WriteString(p.Name)
	if underline {
		b.WriteString(ansi.Reset)
	}
	for range namelen - len(p.Name) {
		b.WriteString(" ")
	}

	b.WriteString("\t")

	// Description
	if underline {
		b.WriteString(ansi.Underline)
	}
	b.WriteString(p.Desc)
	if underline || color != noColor {
		b.WriteString(ansi.Reset)
	}
	if _, err := out.Write([]byte(b.String())); err != nil {
		return err
	}
	return nil
}

func (l ListLints) Run(ctx context.Context) error {
	l.Init()
	if l.OutputJson {
		return l.RunJson(ctx)
	}
	return l.RunTxt(ctx)
}

func (l ListLints) RunTxt(ctx context.Context) error {
	list := make([]prb, 0)
	header := prb{Name: "Name", Desc: "Description", EnabledString: "Status"}
	namelen, enabledlen := len(header.Name), len(header.EnabledString)
	for _, pf := range lints.Available() {
		item := prb{Name: pf.Name(), Desc: pf.Description(), Enabled: pf.Enabled(), EnabledString: enabledString(pf.Enabled())}
		list = append(list, item)
		if len(item.Name) > namelen {
			namelen = len(item.Name)
		}
		if len(item.EnabledString) > enabledlen {
			enabledlen = len(item.EnabledString)
		}
	}
	b := strings.Builder{}
	if err := header.Write(os.Stdout, namelen, enabledlen, &b, l.Color, noColor); err != nil {
		return err
	}

	for _, item := range list {
		c := noColor
		if l.Color {
			if item.Enabled {
				c = green
			} else {
				c = red
			}
		}
		if err := item.Write(os.Stdout, namelen, enabledlen, &b, false, c); err != nil {
			return err
		}

	}
	return nil
}

func (l ListLints) RunJson(ctx context.Context) error {
	for _, pf := range lints.Available() {
		if err := json.MarshalWrite(os.Stdout,
			prb{Name: pf.Name(), Desc: pf.Description(), Enabled: pf.Enabled()},
			jsontext.WithIndent("  ")); err != nil {
			return err
		}
		if _, err := os.Stdout.WriteString("\n"); err != nil {
			return err
		}

	}
	return nil
}

func RunFromCommand(ctx context.Context, command *cli.Command) error {
	return FromCommand(command).Run(ctx)
}
