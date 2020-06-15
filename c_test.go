package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestCLang(t *testing.T) {
	f := foo()
	fmt.Println(f)
}

func TestInt(t *testing.T) {
	printInt(3)
}

// go 에서는 string을 []byte 로 캐스팅하거나 []byte를 string 으로 캐스팅 시에는 복사 발생(string은 불변)
// 복사가 되지 않게 하려면..
//
// C 구조에서는 다음과 같이 정의되어 있음.
// typedef struct { const char *p; GoInt n; } GoString;
// typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
// 그러므로 포인터 만 가져와서 대입 시에는 복사가 일어나지 않음
// 런타임에서 체크는 하고있으므로 대입받은 []byte 의 변경 시도 시에는 panic 이 일어남
// string 에서 []byte 의 동작은 다음과 같이 수행하고 []byte 에서 string 의 변경은
// strings.Builder를 참고할 것.
func TestStringToByte(t *testing.T) {
	s := "hello"
	b1 := []byte(s)
	b2 := *((*[]byte)(unsafe.Pointer(&s)))

	// b2[0] = 'k' // <- panic!
	// b1[0] = 'k' // OK
	fmt.Println(b1)
	fmt.Println(b2)
}

func byteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestStringBuilder(t *testing.T) {
	// byte 배열을 string 처럼 취급하고 있으므로
	// 런타임중에 string 을 수정가능.
	// slice 재할당이 일어날때는 수정 불가..
	b := []byte{104, 101, 108, 108, 111}
	s := byteToString(b)
	fmt.Println(s)
	b[0] = 101
	fmt.Println(s)
}
