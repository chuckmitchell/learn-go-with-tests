package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	walkWithSeen(reflect.ValueOf(x), fn, make(map[uintptr]struct{}))
}

func walkWithSeen(val reflect.Value, fn func(input string), seen map[uintptr]struct{}) {
	walkValue := func(value reflect.Value) {
		walkWithSeen(value, fn, seen)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkWithSeen(val.MapIndex(key), fn, seen)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkWithSeen(v, fn, seen)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkWithSeen(res, fn, seen)
		}
	case reflect.Ptr:
		if !val.IsNil() {
			ptrAddr := val.Pointer() // address the pointer points to
			if _, ok := seen[ptrAddr]; ok {
				return // cycle detected
			}
			seen[ptrAddr] = struct{}{}
			defer delete(seen, ptrAddr)
			walkValue(val.Elem())
		}
	}
}
