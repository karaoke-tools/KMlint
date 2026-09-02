// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package warning

import "github.com/karaoke-tools/kmlint/internal/karajson"

var (
	// warnings
	R18Lyrics = karajson.MustParseTid("e2b8419f-1d5a-44ad-a62c-d7765493190d")
	R18Media  = karajson.MustParseTid("e82ce681-6d7b-4fb6-abe4-daa8aaa9bbf9")
	Spoiler   = karajson.MustParseTid("24371984-5e4c-4485-a937-fb0c480ca23b")
	Epilepsy  = karajson.MustParseTid("51288600-29e0-4e41-a42b-77f0498e5691")
)
