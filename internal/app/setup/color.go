// Copyright Louis Royer. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package setup

import (
	"os"

	"github.com/moby/term"
	"github.com/sirupsen/logrus"
)

func hasNoColorEnv() bool {
	return os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb"
}

func StderrColorAuto() bool {
	return term.IsTerminal(os.Stderr.Fd()) && !hasNoColorEnv()
}

func StdoutColorAuto() bool {
	return term.IsTerminal(os.Stdout.Fd()) && !hasNoColorEnv()
}

func SetLogrusFormatter(color bool) {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:          !color,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
}
