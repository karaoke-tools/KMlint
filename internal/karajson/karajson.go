// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package karajson

import (
	"context"
	"encoding/json/v2"
	"os"
	"time"
)

func FromFile(ctx context.Context, path string) (*KaraJson, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	kara := new(KaraJson)
	if err := json.Unmarshal(content, kara); err != nil {
		return nil, err
	}
	return kara, nil
}

type KaraJson struct {
	Header Header  `json:"header"`
	Medias []Media `json:"medias"`
	Data   Data    `json:"data"`
}

type Header struct {
	Version     int    `json:"version"`
	Description string `json:"description"`
}

type Media struct {
	Default  bool    `json:"default"`
	Duration int16   `json:"duration"`
	Filename string  `json:"filename"`
	Filesize uint64  `json:"filesize"`
	Loudnorm string  `json:"loudnorm"`
	Lyrics   []Lyric `json:"lyrics"`
	Version  string  `json:"version"`
}

type Lyric struct {
	Default  bool   `json:"default"`
	Filename string `json:"filename"`
	Version  string `json:"version"`
}

type Data struct {
	CreatedAt             time.Time         `json:"created_at"`
	FromDisplayType       string            `json:"from_display_type"`
	IgnoreHooks           bool              `json:"ignore-hooks"`
	Kid                   Kid               `json:"kid"`
	ModifiedAt            time.Time         `json:"modified_at"`
	Parents               []Kid             `json:"parents"`
	Repository            string            `json:"repository"`
	Songname              string            `json:"songname"`
	Songorder             *int              `json:"songorder,omitempty"`
	Tags                  Tags              `json:"tags"`
	Titles                map[string]string `json:"titles"`
	TitlesAliases         []string          `json:"titles_aliases"`
	TitlesDefaultLanguage string            `json:"titles_default_language"`
	Year                  int               `json:"year"`
}

type Tags struct {
	Authors      []Tid `json:"authors"`
	Collections  []Tid `json:"collections"`
	Creators     []Tid `json:"creators"`
	Families     []Tid `json:"families"`
	Groups       []Tid `json:"groups"`
	Langs        []Tid `json:"langs"`
	Misc         []Tid `json:"misc"`
	Origins      []Tid `json:"origins"`
	Platforms    []Tid `json:"platforms"`
	Series       []Tid `json:"series"`
	Singers      []Tid `json:"singers"`
	Singergroups []Tid `json:"singergroups"`
	Songtypes    []Tid `json:"songtypes"`
	Songwriters  []Tid `json:"songwriters"`
	Versions     []Tid `json:"versions"`
	Warnings     []Tid `json:"warnings"`
}
