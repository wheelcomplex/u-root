// Copyright 2012 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// our first builtin: cd
package main

import (
	"errors"
	"os"
)

func init() {
	addBuiltIn("cd", cd)
}

func cd(s []string) error {
	if len(s) != 1 {
		return errors.New("usage: cd one-path")
	}

	err := os.Chdir(s[0])
	return err
}
