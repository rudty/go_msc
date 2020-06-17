package cache

import (
	"fmt"
	"reflect"
)

type Any interface{}

type leap struct {
	value []reflect.Value
	next  map[Any]Any
}

type FunctionCache struct {
	leap     *leap
	Function Any
}

func NewFunctionCache(fn Any) *FunctionCache {
	f := &FunctionCache{}
	f.Function = fn
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
		cNode.value = reflect.ValueOf(f.Function).Call(fArgs)
	}
	return cNode.value
}

func (f *FunctionCache) Call(args ...Any) Any {
	ret := f.call(args...)
	return ret[0].Interface()
}

func (f *FunctionCache) CallR2(args ...Any) (Any, Any) {
	ret := f.call(args...)
	if len(ret) != 2 {
		panic(fmt.Sprint("function result length != 2. result: ", len(ret)))
	}
	r0 := ret[0].Interface()
	r1 := ret[1].Interface()
	return r0, r1
}

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
