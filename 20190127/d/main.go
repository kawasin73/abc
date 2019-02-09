package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := nextInt()

	childs := make([][]int, n)

	all := n + m - 1
	for i := 0; i < all; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		childs[b] = append(childs[b], a)
	}

	depth := make([]int, n)
	for i := 0; i < n; i++ {
		depth[i] = -1
	}
	parents := make([]int, n)

	for i := 0; i < n; i++ {
		ps := len(childs[i])
		if ps == 0 {
			// root
			parents[i] = -1
		} else if ps == 1 {
			parents[i] = childs[i][0]
		} else {
			maxdepth := -1
			maxid := -1
			for _, a := range childs[i] {
				depth := findDepth(depth, childs, a)
				if depth > maxdepth {
					maxdepth = depth
					maxid = a
				}
			}
			parents[i] = maxid
		}
	}

	//fmt.Println(childs)
	//fmt.Println(parents)
	//fmt.Println(depth)
	for _, p := range parents {
		fmt.Println(p + 1)
	}
}

func findDepth(depth []int, childs [][]int, i int) int {
	if depth[i] != -1 {
		return depth[i]
	}
	maxdepth := -1
	for _, a := range childs[i] {
		depth := findDepth(depth, childs, a)
		if depth > maxdepth {
			maxdepth = depth
		}
	}
	maxdepth++
	depth[i] = maxdepth
	return maxdepth
}
