package cache

import (
	"reflect"
)

type leap struct {
	value interface{}
	next  map[interface{}]interface{}
}

type FunctionCache struct {
	leap     leap
	Function interface{}
}

func NewFunctionCache(fn interface{}) *FunctionCache {
	f := &FunctionCache{}
	f.Function = fn
	f.leap.next = make(map[interface{}]interface{})
	return f
}

func (f *FunctionCache) Call(args ...interface{}) interface{} {
	a := f.leap
	m := a.next
	for i := 0; i < len(args); i++ {
		value, ok := m[args[i]]
		if !ok {
			newLeap := &leap{
				value: nil,
				next:  make(map[interface{}]interface{}),
			}
			m[args[i]] = newLeap
			value = newLeap
		}
		a = (value.(leap))
	}

	if a.value == nil {
		fArgs := make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			fArgs[i] = reflect.ValueOf(args[i])
		}
		a.value = reflect.ValueOf(f.Function).Call(fArgs)
	}
	return a.value
}
