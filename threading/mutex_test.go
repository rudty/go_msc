package threading

import (
	"testing"
	"time"
)

func TestCreateMutex(t *testing.T) {
	m := NewMutex()

	sum := 0
	for _, e := range m.data {
		sum += int(e)
	}

	if sum == 0 {
		t.Error("mutex init fail")
	}
}

func TestReleaseMutex(t *testing.T) {
	m := NewMutex()
	m.Release()
	sum := 0
	for _, e := range m.data {
		sum += int(e)
	}
	if sum != 0 {
		t.Error("mutex release fail")
	}
}

var testLockV = 0
var testLockMutex *Mutex

func testLockIncWithLock() {
	for i := 0; i < 1000000; i++ {
		testLockMutex.Lock()
		testLockV++
		testLockMutex.Unlock()
	}
}
func TestLock(t *testing.T) {
	testLockMutex = NewMutex()

	t1 := NewThread(testLockIncWithLock)
	t2 := NewThread(testLockIncWithLock)
	t1.Start()
	t2.Start()

	t1.Join()
	t2.Join()

	if testLockV != 2000000 {
		t.Error("mutex lock error")
	}
}

var testIDUniqueMap = make(map[uint64]bool)
var testIDUniqueMutex *Mutex

func testIDUniqueRegisterID() {
	time.Sleep(1 * time.Second)
	testIDUniqueMutex.Lock()
	testIDUniqueMap[GetThreadID()] = true
	testIDUniqueMutex.Unlock()
}

func TestTIdUnique(t *testing.T) {
	testIDUniqueMutex = NewMutex()

	threads := make([]*Thread, 1000)

	for i := 0; i < 1000; i++ {
		threads[i] = NewThread(testIDUniqueRegisterID)
		threads[i].Start()
	}

	for i := 0; i < 1000; i++ {
		threads[i].Join()
	}

	if len(testIDUniqueMap) != 1000 {
		t.Error("thread count = 1000")
	}
}
