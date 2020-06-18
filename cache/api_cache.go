package cache

import (
	"fmt"
	"reflect"
)

// Any 함수 반환형식
type Any interface{}

type leap struct {
	value []reflect.Value
	next  map[Any]Any
}

// FunctionCache 함수를 캐싱할 수 있는 구조입니다.
// 한번 호출한 함수는 내부적으로 저장해서 다음번 호출 시에는
// 저장한 반환값을 다시 반환합니다.
type FunctionCache struct {
	leap *leap
	fn   Any
}

// NewFunctionCache 새로운 캐싱할 수 있는 함수를 만듭니다.
func NewFunctionCache(fn Any) *FunctionCache {
	f := &FunctionCache{}
	f.fn = fn
	f.leap = &leap{
		value: nil,
		next:  make(map[Any]Any),
	}
	return f
}

func (f *FunctionCache) call(args ...Any) []reflect.Value {
	cNode := f.leap
	argMap := cNode.next
	for i := 0; i < len(args); i++ {
		value, ok := argMap[args[i]]
		if !ok {
			newLeap := &leap{
				value: nil,
				next:  make(map[Any]Any),
			}
			argMap[args[i]] = newLeap
			value = newLeap
		}
		cNode = (value.(*leap))
	}

	if cNode.value == nil {
		fArgs := make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			fArgs[i] = reflect.ValueOf(args[i])
		}
		cNode.value = reflect.ValueOf(f.fn).Call(fArgs)
	}
	return cNode.value
}

// Call 함수를 호출하고 첫번째 반환값을 가져옵니다.
// 반환값은 반드시 1개 이상이 있어야 합니다
// 내부적으로 캐싱이 되어있다면 캐싱된 반환값을 가져옵니다
func (f *FunctionCache) Call(args ...Any) Any {
	ret := f.call(args...)
	return ret[0].Interface()
}

// CallR2 함수를 호출하고 2개의 반환값을 가져옵니다.
// 반환값은 반드시 2개 이상이 있어야 합니다
// 내부적으로 캐싱이 되어있다면 캐싱된 반환값을 가져옵니다
func (f *FunctionCache) CallR2(args ...Any) (Any, Any) {
	ret := f.call(args...)
	if len(ret) != 2 {
		panic(fmt.Sprint("function result length != 2. result: ", len(ret)))
	}
	r0 := ret[0].Interface()
	r1 := ret[1].Interface()
	return r0, r1
}

// CallR3 함수를 호출하고 3개의 반환값을 가져옵니다.
// 반환값은 반드시 2개 이상이 있어야 합니다
// 내부적으로 캐싱이 되어있다면 캐싱된 반환값을 가져옵니
func (f *FunctionCache) CallR3(args ...Any) (Any, Any, Any) {
	ret := f.call(args...)
	if len(ret) != 3 {
		panic(fmt.Sprint("function result length != 3. result: ", len(ret)))
	}
	r0 := ret[0].Interface()
	r1 := ret[1].Interface()
	r2 := ret[2].Interface()
	return r0, r1, r2
}
