package mymap

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1)
	t.Log(m1[0], m1[1], m1[2], m1[3])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m2=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key s is %d", v)
	} else {
		t.Log("key s is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](1), m[2](2), m[3](3))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(mySet, 1)

	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}

func TestMapfuncvalue(t *testing.T) {
	Mapfuncvalue()
}

func TestMapvaluenil(t *testing.T) {
	Mapvaluenil()
}

func TestMapvaluenil1(t *testing.T) {
	Mapvaluenil1()
}

func TestNonConcurrentMap(t *testing.T) {

	// fatal error: concurrent map read and map write
	m := make(map[int]int)

	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {

	}
}

func TestConcurrentMap(t *testing.T) {
	var scene sync.Map
	i := 0
	for i < 1000 {
		scene.Store("creece"+strconv.Itoa(i), i)
		// scene.Store("creece", 98)
		// scene.Store("egypt", 200)
		// scene.Store("London", 201)

		// if count, ok := scene.Load("London"); ok {
		// 	newCount := fmt.Sprintf("%d", count)
		// 	counts, _ := strconv.Atoi(newCount)
		// 	t.Log(ok, counts-1, reflect.TypeOf(count))
		// }
		counts := ParseSyncMap(&scene, "London")
		if counts > 0 {
			t.Log(counts)
		}
		// scene.Delete("London")
		// scene.Store("London", 201)
		scene.Range(func(k, v interface{}) bool {
			fmt.Printf("iterator: %v %v\n", k, v)
			return true
		})
		i++
	}

}

func ParseSyncMap(p *sync.Map, key string) int {
	if count, ok := p.Load(key); ok {
		newCount := fmt.Sprintf("%d", count)
		counts, _ := strconv.Atoi(newCount)
		return counts
	}
	return 0
}
