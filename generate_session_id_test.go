package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

var index uint32 = 998
var machineUniqueValue []byte

// NewSessionID 새로운 세션 아이디를 만듭니다
func NewSessionID() string {
	// b := make([]byte, 12+len(machineUniqueValue))
	now := time.Now().Unix()
	newIndex := atomic.AddUint32(&index, 1)
	// binary.LittleEndian.PutUint64(b[:], uint64(now))
	// binary.LittleEndian.PutUint32(b[8:], newIndex)

	// copy(b[12:], machineUniqueValue)
	// fmt.Println(b)
	// fmt.Println(hex.EncodeToString(b))
	// r := base64.StdEncoding.EncodeToString(b)
	// fmt.Println(r)
	// k, _ := base64.StdEncoding.DecodeString(r)
	// fmt.Println(k)

	return fmt.Sprintf("%d%03d%s",
		now,
		newIndex%1000,
		defaultMachineUniqueValue2())
}

func defaultMachineUniqueValue() []byte {
	var b [12]byte
	pid := uint64(os.Getpid())
	randomValue := rand.Uint32()
	binary.LittleEndian.PutUint64(b[:], pid)
	binary.LittleEndian.PutUint32(b[8:], randomValue)
	return b[:]
}

func defaultMachineUniqueValue2() string {
	pid := os.Getpid()
	randomValue := rand.Uint32() % 10
	return fmt.Sprintf("%d%d", pid, randomValue)
}

func init() {
	machineUniqueValue = defaultMachineUniqueValue()
}

func TestSessionID(t *testing.T) {
	fmt.Println(NewSessionID())
}
