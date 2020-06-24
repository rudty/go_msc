// +build windows

package threading

import "syscall"

// GetThreadID get thread id
func GetThreadID() uint64 {
	r0, _, e1 := syscall.Syscall(procGetCurrentThreadID.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		return 0
	}
	return uint64(r0)
}
