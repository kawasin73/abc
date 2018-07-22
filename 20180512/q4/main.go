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
	sc.Split(bufio.ScanWords)

	N := nextInt()
	M := nextInt()

	ps := make(map[int]int, N)
	for i := 0; i < N; i++ {
		ps[nextInt()] = i + 1
	}

	memo := make(map[int]map[int]bool)
	for i := 0; i < M; i++ {
		x := nextInt()
		y := nextInt()
		src, ok := memo[y]
		if !ok {
			src = make(map[int]bool)
			memo[y] = src
		}
		src[x] = true

		src, ok = memo[x]
		if !ok {
			src = make(map[int]bool)
			memo[x] = src
		}
		src[y] = true
	}

	var result int
	for i := N; i > 0; i-- {
		if ps[i] == i {
			result++
			continue
		}
		fmt.Println("find ", i)
		if find(ps[i], i, memo) {
			result++
		}
	}
	fmt.Println(result)
}

func main2() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	M := nextInt()

	ps := make([]int, N)
	for i := 0; i < N; i++ {
		ps[i] = nextInt()
	}

	memo := make([]int, N)
	for i := 0; i < M; i++ {
		x := nextInt()
		y := nextInt()

		if memo[x] != 0 {
			memo[y] = memo[x]
		} else if memo[y] != 0 {
			memo[x] = memo[y]
		} else {
			memo[x] = y
		}
	}


	var result int
	for i := N; i > 0; i-- {
		if ps[i] == i {
			result++
			continue
		}
		fmt.Println("find ", i)
		if find(ps[i], i, memo) {
			result++
		}
	}
	fmt.Println(result)
}

func find(from, to int, memo map[int]map[int]bool) bool {
	fmt.Println("debug: ", from, to)
	m, ok := memo[to]
	if !ok {
		return false
	}
	if result, ok := m[from]; ok {
		return result
	}
	for f, result := range m {
		if result {
			if find(from, f, memo) {
				m[from] = true
				return true
			}
		}
	}
	m[from] = false
	return false
}
