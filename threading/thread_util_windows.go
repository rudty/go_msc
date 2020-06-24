// +build windows

package threading

import "syscall"

var modkernel32 = syscall.NewLazyDLL("kernel32.dll")
var (
	procGetCurrentThreadID        = modkernel32.NewProc("GetCurrentThreadId")
	procInitializeCriticalSection = modkernel32.NewProc("InitializeCriticalSection")
	procDeleteCriticalSection     = modkernel32.NewProc("DeleteCriticalSection")
	procEnterCriticalSection      = modkernel32.NewProc("EnterCriticalSection")
	procLeaveCriticalSection      = modkernel32.NewProc("LeaveCriticalSection")
)
