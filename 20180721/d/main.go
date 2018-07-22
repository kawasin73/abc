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

type pair struct {
	start int
	end   int
}

func main() {
	sc.Split(bufio.ScanWords)

	// start -> id map
	startMap := make(map[int]map[int]struct{})
	// end -> id map
	endMap := make(map[int]map[int]struct{})
	idMap := make(map[int]pair)

	_ = nextInt()
	m := nextInt()

	for i := 0; i < m; i ++ {
		start := nextInt()
		end := nextInt() - 1
		if _, ok := startMap[start]; ok {
			startMap[start][i] = struct{}{}
		} else {
			startMap[start] = map[int]struct{}{i: {}}
		}

		if _, ok := endMap[end]; ok {
			endMap[end][i] = struct{}{}
		} else {
			endMap[end] = map[int]struct{}{i: {}}
		}

		idMap[i] = pair{start: start, end: end}
	}

	startList := make([]int, len(startMap))
	var i int
	for start, _ := range startMap {
		startList[i] = start
		i++
	}
	sort.Ints(startList)

	endList := make([]int, len(endMap))
	i = 0
	for end, _ := range endMap {
		endList[i] = end
		i++
	}
	sort.Ints(endList)

	var count int
	var currentEnd int
	for _, start := range startList {
		if start <= currentEnd {
			ids := startMap[start]
			for id, _ := range ids {
				delete(idMap, id)
			}
			continue
		}

		var found bool
	endloop:
		for ; len(endList) > 0; {
			end := endList[0]
			endList = endList[1:]
			ids := endMap[end]
			for id, _ := range ids {
				if _, ok := idMap[id]; ok {
					found = true
					currentEnd = end
					break endloop
				}
			}
		}
		if !found {
			break
		}

		ids := startMap[start]
		for id, _ := range ids {
			delete(idMap, id)
		}
		count++
	}

	fmt.Println(count)
}
