package main

import (
	"bufio"
	"os"
	"strconv"
	"sort"
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

	as := make([]int, n)

	for i := 0; i < n; i++ {
		as[i] = nextInt() - i - 1
	}
	sort.IntSlice(as).Sort()
	b := as[n/2]
	var sum int
	for i := 0; i < n; i++ {
		bb := as[i] - b
		if bb > 0 {
			sum += bb
		} else {
			sum -= bb
		}
	}
	fmt.Println(sum)
}
