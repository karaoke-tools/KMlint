// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package ass

import (
	"bufio"
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/karaoke-tools/kmlint/internal/ass/lyrics"
)

type Ass struct {
	Filename                string
	ScriptInfo              ScriptInfo
	Styles                  []string
	Events                  []*lyrics.LyricsParser
	Extradata               []string
	AegisubGarbage          bool // has a Aegisub Project Garbage section
	Fonts                   bool // has Fonts section
	assUnknownSectionsCount int
}

func New(filename string) *Ass {
	return &Ass{
		Filename: filename,
		Styles:   make([]string, 0),
		Events:   make([]*lyrics.LyricsParser, 0),
	}
}

func (ass *Ass) Parse(ctx context.Context) error {
	f, err := os.OpenFile(ass.Filename, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	state := assInit
	i := 0
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			i++
			line := scanner.Text()
			if i == 1 {
				// In some files, BOM is present multiple times for no reason
				BOM := string([]byte{0xEF, 0xBB, 0xBF})
				for strings.HasPrefix(line, BOM) {
					line = strings.TrimPrefix(line, BOM)
				}
			}
			if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
				switch strings.TrimSuffix(strings.TrimPrefix(line, "["), "]") {
				case "Script Info":
					state = assScriptInfo
				case "Aegisub Project Garbage":
					state = assAegisubGarbage
					ass.AegisubGarbage = true
				case "V4+ Styles":
					state = assStyles
				case "Fonts":
					state = assFonts
					ass.Fonts = true
				case "Events":
					state = assEvents
				case "Aegisub Extradata":
					state = assAegisubExtradata
				default:
					state = assUnknownSection
					ass.assUnknownSectionsCount += 1
				}
				continue
			}
			if line == "" {
				// empty line
				continue
			}
			switch state {
			case assScriptInfo:
				if strings.HasPrefix(line, ";") {
					// comment line
					continue
				}
				lineSplit := strings.SplitN(line, ": ", 2)
				if len(lineSplit) != 2 {
					// unreadable
					continue
				}
				switch lineSplit[0] {
				case "PlayResX":
					res, err := strconv.ParseUint(lineSplit[1], 10, 32)
					if err != nil {
						return ErrMalformedFile
					}
					ass.ScriptInfo.PlayResX = uint32(res)
				case "PlayResY":
					res, err := strconv.ParseUint(lineSplit[1], 10, 32)
					if err != nil {
						return ErrMalformedFile
					}
					ass.ScriptInfo.PlayResY = uint32(res)
				case "ScaledBorderAndShadow":
					if err := ass.ScriptInfo.SetScaledBorderAndShadow(lineSplit[1]); err != nil {
						return err
					}
				}
			case assStyles:
				ass.Styles = append(ass.Styles, line)
			case assEvents:
				lyr, err := lyrics.Parse(line)
				if err != nil {
					return err
				}
				ass.Events = append(ass.Events, lyr)
			case assAegisubExtradata:
				ass.Extradata = append(ass.Extradata, line)
			default:
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
