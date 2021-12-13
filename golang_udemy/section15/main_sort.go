package main

import (
	"fmt"
	"sort"
)

// sort

type Entry struct {
	Name  string
	Value int
}

func main() {
	// i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	// s := []string{"a", "z", "j"}

	// fmt.Println(i, s)
	// sort.Ints(i)
	// sort.Strings(s)

	// fmt.Println(i, s)

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
		{"B", 10},
	}

	fmt.Println(el)

	// slice 2回実行すると2回目の結果が反映される
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.Slice(el, func(i, j int) bool { return el[i].Value < el[j].Value })

	fmt.Println(el)

	// SliceStable 安定ソートを保証 同じvalueの場合は、value内でnameの順番にsortされる
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })
	sort.SliceStable(el, func(i, j int) bool { return el[i].Value < el[j].Value })
	fmt.Println(el)
}
