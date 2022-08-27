package main

import (
	"fmt"
	"sort"
)

type sortMap struct {
	m    map[interface{}]interface{}
	keys []interface{}
}

func main() {

	m := map[string]int{"Apple": 20, "Tomato": 12, "Banana": 18}
	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
