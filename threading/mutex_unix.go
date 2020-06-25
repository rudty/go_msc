// +build !windows

package threading

import (
	"runtime"
	"unsafe"
)

/*
#ifndef _GNU_SOURCE
#define _GNU_SOURCE
#endif

#include <pthread.h>
pthread_mutexattr_t recursive_attr;
void init_mutex_attr() {
	pthread_mutexattr_init(&recursive_attr);
	pthread_mutexattr_settype(&recursive_attr, PTHREAD_MUTEX_RECURSIVE);
}
*/
import "C"

func init() {
	C.init_mutex_attr()
}

func (m *Mutex) pthreadMutex() *C.pthread_mutex_t {
	return (*C.pthread_mutex_t)(unsafe.Pointer(m))
}

// NewMutex init mutex call pthread_mutex_init
func NewMutex() *Mutex {
	m := &Mutex{}
	C.pthread_mutex_init(m.pthreadMutex(), &C.recursive_attr)
	runtime.SetFinalizer(m, (*Mutex).pthreadMutexDestory)
	return m
}

func (m *Mutex) pthreadMutexDestory() {
	C.pthread_mutex_destroy(m.pthreadMutex())
	for i := 0; i < len(m.data); i++ {
		m.data[i] = 0
	}
}

// Release mutex call pthread_mutex_destroy
func (m *Mutex) Release() {
	runtime.SetFinalizer(m, nil)
	m.pthreadMutexDestory()
}

// Lock mutex call pthread_mutex_lock
func (m *Mutex) Lock() {
	C.pthread_mutex_lock(m.pthreadMutex())
}

// Unlock mutex call pthread_mutex_unlock
func (m *Mutex) Unlock() {
	C.pthread_mutex_unlock(m.pthreadMutex())
}
