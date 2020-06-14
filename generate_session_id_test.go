package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"testing"
	"time"
)

var index uint32 = 998
var machineUniqueValue string

// NewSessionID 새로운 세션 아이디를 만듭니다
func NewSessionID() string {
	now := time.Now().Unix()
	newIndex := atomic.AddUint32(&index, 1)

	return fmt.Sprintf("%03d%d%s",
		newIndex%1000,
		now,
		machineUniqueValue)
}

// defaultMachineUniqueValue 각 서버별 고유 아이디
// pid + 랜덤이면 소형 서버에서는 고유할것으로 생각
func defaultMachineUniqueValue() string {
	pid := os.Getpid()
	randomValue := rand.Uint32() % 100
	return fmt.Sprintf("%d%02d", pid, randomValue)
}

func init() {
	machineUniqueValue = defaultMachineUniqueValue()
}

func TestSessionID(t *testing.T) {
	fmt.Println(NewSessionID())
}
