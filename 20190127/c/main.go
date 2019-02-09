package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

type item struct {
	i   int
	sum int
}

type items []item

func (it *items) Len() int {
	return len(*it)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (it *items) Less(i, j int) bool {
	return (*it)[i].sum > (*it)[j].sum
}

// Swap swaps the elements with indexes i and j.
func (it *items) Swap(i, j int) {
	(*it)[i], (*it)[j] = (*it)[j], (*it)[i]
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	buf := make([]item, n)
	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
		buf[i].sum = a[i] + b[i]
		buf[i].i = i
	}

	is := items(buf)
	sort.Sort(&is)

	var sum int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += a[buf[i].i]
		} else {
			sum -= b[buf[i].i]
		}
	}
	fmt.Println(sum)
}
