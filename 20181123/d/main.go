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

var a = make([]int, 60)

func main() {
	sc.Split(bufio.ScanWords)
	for i := 2; i <= 30; i++ {
		a[i] = nextInt()
	}

	if a[2] > 41 {
		fmt.Println("invalid")
		return
	}

	r := createInt(0, 0, a[2])
	if r != -1 {
		fmt.Println(r)
	} else {
		fmt.Println("invalid")
	}
}

func createInt(test uint, head uint, rest int) int {
	if rest == 0 {
		if check(int(test)) {
			return int(test)
		}
		return -1
	} else if head > 41 {
		return -1
	}
	if r := createInt(test, head+1, rest); r != -1 {
		return r
	}
	test |= 1 << head
	if r := createInt(test, head+1, rest-1); r != -1 {
		return r
	}
	return -1
}

func check(test int) bool {
	for i := 3; i <= 30; i++ {
		if a[i] != sum(test, i) {
			return false
		}
	}
	return true
}

func sum(test int, i int) int {
	var s int
	for test > 0 {
		s += test % i
		test /= i
	}
	return s
}
