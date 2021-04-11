// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux
// +build js

package unix

const (
	getrandomTrap     uintptr = 0
	copyFileRangeTrap uintptr = 0
)
