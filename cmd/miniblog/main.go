// Copyright 2022 Innkeeper Belm(卢明健) &lt;2814104062@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is git@github.com:lmj-username/lmj-username.git

package main

import (
	"miniblog/internal/miniblog"
	"os"
)

func main() {
	command := miniblog.NewMiniBlogCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
