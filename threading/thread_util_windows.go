// +build windows

package threading

import "syscall"

var modkernel32 = syscall.NewLazyDLL("kernel32.dll")
var procGetCurrentThreadID = modkernel32.NewProc("GetCurrentThreadId")
