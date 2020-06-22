// +build !windows

package threading

/*
#include <pthread.h>
unsigned long long getThreadId() {
	unsigned long long tid;
	pthread_threadid_np(NULL, &tid);
	return tid;
}
*/
import "C"

// GetThreadID get thread id
func GetThreadID() uint64 {
	return uint64(C.getThreadId())
}
