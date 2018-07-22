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
	n := nextInt()

	index := []int{1, 6, 9, 36, 81, 216, 729, 1296, 6561, 7776, 46656, 59049, 279936, 531441}

	var count int
	for len(index) > 0 {
		last := index[len(index)-1]
		if last == 7776 {
			break
		}
		if n < last {
			index = index[:len(index)-1]
			continue
		}
		n -= index[len(index)-1]
		count++
	}

	hoge := index[:len(index)-1]
	c2 := search(n, count, hoge)
	for n >= 7776 {
		n -= 7776
		count++
		c := search(n, count, hoge)
		if c2 == 0 || c < c2 {
			c2 = c
		}
	}


	fmt.Println(c2)
}

func search(n int, count int, index []int) int {
	for len(index) > 0 {
		last := index[len(index)-1]
		if n < last {
			index = index[:len(index)-1]
			continue
		}
		n -= last
		count++
	}
	return count
}
