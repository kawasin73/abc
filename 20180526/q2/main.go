package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	if !sc.Scan() {
		panic("failed scan string")
	}
	return sc.Text()
}

func nextInt() int {
	if !sc.Scan() {
		panic("failed scan int")
	}
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	n := nextInt()
	s := []byte(nextString())

	max := 0
	for i := 0; i < n; i++ {
		k := count(i, s)
		if k > max {
			max = k
		}
	}
	fmt.Println(max)
}

func count(i int, s []byte) int {
	s1 := s[:i]
	s2 := s[i:]

	m1 := map[byte]bool{}
	m2 := map[byte]bool{}

	for _, k := range s1 {
		m1[k] = true
	}
	for _, k := range s2 {
		m2[k] = true
	}

	var c int
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			c++
		}
	}

	return c
}
