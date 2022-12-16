package main

import (
	"fmt"
	"strconv"
)

const (
	things = iota
	persons
	animals
	long
	thinAndFlat
)

func Counter(t int) string {
	s := ""
	m := make(map[int]string)
	m[things] = "abc"
	m[persons] = "def"
	m[animals] = "ghi"
	m[long] = "hij"
	m[thinAndFlat] = "lmn"

	for i := 1; i <= 100; i++ {
		s += fmt.Sprintf("%v\n", m[t]+strconv.Itoa(i))
	}

	return s
}
