package syscall

import (
	_ "unsafe" // for go:linkname
)

//go:linkname runtime_envs runtime.get_envs
func runtime_envs() []string
