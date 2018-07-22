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
	as := make([]int, n)

	for i := 0; i < n; i++ {
		as[i] = nextInt()
	}

	var sums [4]int
	index := [4]int{0, 1, 2, 3}

	sums[0] = as[0]
	sums[1] = as[1]
	sums[2] = as[2]
	for i := 3; i < n; i++ {
		sums[3] += as[i]
	}
	// min -> max
	order := []int{0, 1, 2, 3}

	sortOrder := func() {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if sums[order[j]] > sums[order[j+1]] {
					order[j], order[j+1] = order[j+1], order[j]
				}
			}
		}
	}
	sortOrder()

	move := func(i int) bool {
		// min
		if i == order[0] {
			return false
		}
		// max
		if i-1 == order[3] {
			return false
		}
		delSum := sums[i] - as[index[i]]
		addSum := sums[i-1] + as[index[i]]

		if delSum < sums[order[0]] {
			return false
		}
		if addSum > sums[order[3]] {
			return false
		}
		sums[i] = delSum
		sums[i-1] = addSum
		index[i]++
		sortOrder()
		return true
	}

	for {
		if index[3]+1 < n && move(3) {
			continue
		} else if index[2]+1 < index[3] && move(2) {
			continue
		} else if index[1]+1 < index[2] && move(1) {
			continue
		} else {
			break
		}
	}
	fmt.Println(sums[order[3]] - sums[order[0]])
}
