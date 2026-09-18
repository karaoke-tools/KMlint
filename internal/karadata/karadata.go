// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package karadata

import (
	"context"
	"path"

	"github.com/karaoke-tools/kmlint/internal/ass"
	"github.com/karaoke-tools/kmlint/internal/karajson"
)

// Song information
type KaraData struct {
	KaraJson karajson.KaraJson // metadata of the song
	Lyrics   []*ass.Ass        // lyrics of the song
}

// Create a new `KaraData` from a `KaraJson`
func New(basedir string, karaJson karajson.KaraJson) (KaraData, error) {
	if len(karaJson.Medias) == 0 {
		return KaraData{}, ErrNoMedias
	}
	data := KaraData{
		KaraJson: karaJson,
		Lyrics:   make([]*ass.Ass, 0, len(karaJson.Medias[0].Lyrics)),
	}
	// TODO: update this when multi-track drifting is released
	for _, l := range data.KaraJson.Medias[0].Lyrics {
		lyricsPath := path.Join(basedir, "lyrics", l.Filename)

		lyrics := ass.New(lyricsPath)
		data.Lyrics = append(data.Lyrics, lyrics)
	}
	return data, nil
}

func (k *KaraData) ParseLyrics(ctx context.Context) error {
	var err error
	for _, l := range k.Lyrics {
		err = l.Parse(ctx)
		if err != nil {
			return err
		}
	}
	return err
}
