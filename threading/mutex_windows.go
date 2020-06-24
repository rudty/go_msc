// +build windows

package threading

import (
	"runtime"
	"syscall"
	"unsafe"
)

// NewMutex init mutex call InitializeCriticalSection
func NewMutex() *Mutex {
	m := &Mutex{}
	_, _, e1 := syscall.Syscall(procInitializeCriticalSection.Addr(), 1, uintptr(unsafe.Pointer(m)), 0, 0)
	if e1 != 0 {
		return nil
	}

	runtime.SetFinalizer(m, (*Mutex).deleteCriticalSection)
	return m
}

func (m *Mutex) deleteCriticalSection() {
	syscall.Syscall(procDeleteCriticalSection.Addr(), 1, uintptr(unsafe.Pointer(m)), 0, 0)
}

// Release mutex call DeleteCriticalSection
func (m *Mutex) Release() {
	runtime.SetFinalizer(m, nil)
	m.deleteCriticalSection()
}

// Lock mutex call EnterCriticalSection
func (m *Mutex) Lock() {
	syscall.Syscall(procEnterCriticalSection.Addr(), 1, uintptr(unsafe.Pointer(m)), 0, 0)
}

// Unlock mutex call procLeaveCriticalSection
func (m *Mutex) Unlock() {
	syscall.Syscall(procLeaveCriticalSection.Addr(), 1, uintptr(unsafe.Pointer(m)), 0, 0)
}
