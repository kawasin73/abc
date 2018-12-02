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

const max = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	var r int
	for i := 1; i <= n; i++ {
		for j := 1; j <= n/i; j++ {
			r += pair(i, j)
			r %= max
		}
	}

	fmt.Println(r)
}

func pair(a, b int) int {
	r1 := 1
	r2 := 1
	r3 := 1
	r4 := 1
	for i := 0; i < 10; i++ {
		r1 *= a
		r2 *= a - 1
		r3 *= b
		r4 *= b - 1

		r1 %= max
		r2 %= max
		r3 %= max
		r4 %= max
	}

	k1 := (r1 - r2 + max) % max
	k2 := (r3 - r4 + max) % max
	return (k1 * k2) % max
}
