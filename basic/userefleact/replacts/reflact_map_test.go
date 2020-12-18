package replacts

import (
	"reflect"
	"testing"
	"time"
)

func cmpMap(m1, m2 map[string]int) bool {
	for k1, v1 := range m1 {
		if v2, has := m2[k1]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, has := m2[k2]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func TestCmpMap(t *testing.T) {
	start := time.Now()

	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < 100000; i++ {
		res := cmpMap(m1, m2)
		_ = res
	}
	end := time.Now()
	du := end.Sub(start)
	t.Log("100000 call cmpMap(m1,m2) elapsed=", du)
}

func TestReflactDeepEqual(t *testing.T) {
	start := time.Now()
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < 100000; i++ {
		res := reflect.DeepEqual(m1, m2)
		_ = res
	}
	end := time.Now()
	du := end.Sub(start)
	t.Log("100000 call reflact.DeepEqual(m1,m2) elapsed=", du)
}

// 由上可知，当不知道类型时可用反射，免去繁琐的判断，否则重写比较函数
