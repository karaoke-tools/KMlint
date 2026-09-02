// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package origin

import "github.com/karaoke-tools/kmlint/internal/karajson"

var (
	// origins
	Fanworks                 = karajson.MustParseTid("ac5085db-52d8-44b4-8fbb-1697b24d2046")
	Movie                    = karajson.MustParseTid("c8ca9dea-d8b4-48fe-8427-d911653232e9")
	Musical                  = karajson.MustParseTid("75b8f183-59d6-4294-bc2e-d1e0f34a1443")
	OriginalNetworkAnimation = karajson.MustParseTid("8db3e098-06ef-45e5-bba9-cb04a21c1e66")
	OriginalVideoAnimation   = karajson.MustParseTid("2a6ef087-a000-4f8e-aeb5-9a7968756a36")
	Show                     = karajson.MustParseTid("8dd8d2a4-913a-4b0b-a555-1121cc603a57")
	TVSpecial                = karajson.MustParseTid("65f74288-49a5-47b0-82f5-8cd84d1e3dc0")
	TVSeries                 = karajson.MustParseTid("938de218-5343-4865-94d3-fb33f2eaa152")
	VideoGame                = karajson.MustParseTid("dbedd6b3-d125-4cd8-aa32-c4175e4ca3a3")
	VisualNovel              = karajson.MustParseTid("2f96e82c-c61b-4413-817e-e6b247509b95")
	Vtuber                   = karajson.MustParseTid("5106d71b-a78e-4d19-a16d-f8fb6719538a")

	// origins alias
	ONA = OriginalNetworkAnimation
	OVA = OriginalVideoAnimation
)
