package main

import (
	"bufio"
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

type items []*line

func (it *items) Len() int {
	return len(*it)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (it *items) Less(i, j int) bool {
	return (*it)[i].y > (*it)[j].y
}

// Swap swaps the elements with indexes i and j.
func (it *items) Swap(i, j int) {
	(*it)[i], (*it)[j] = (*it)[j], (*it)[i]
}

type line struct {
	a      int
	b      int
	y      int
	action struct {
		changed  bool
		from int
	}
	dead bool
}

type node struct {
	root int
	parents []*line
	childs []*line
	sum int
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := nextInt()

	xs := make([]int, n)
	lines := make([]*line, m)
	nodes := make([]node, n)

	for i := 0; i < n; i++ {
		xs[i] = nextInt()
		nodes[i].sum = xs[i]
	}

	for i := 0; i < m; i++ {
		a := nextInt()
		b := nextInt()
		y := nextInt()
		if a > b {
			a, b = b, a
		}
		lines[i] = &line{
			a: a,
			b: b,
			y: y,
		}
	}

	hoge := items(lines)
	sort.Sort(&hoge)

	groups := make([]int, n)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		groups[i] = i
	}

	// union
	for i := m - 1; i >= 0; i-- {
		a := lines[i].a
		b := lines[i].b
		aroot := root(groups, a)
		if groups[b] != aroot {
			lines[i].action.changed = true
			lines[i].action.from = groups[b]
			groups[b] = aroot
			sum[aroot] += sum[b]
		}
	}

	var deleted int
	for i := 0; i < m; i++ {
		b := lines[i].b
		rt := root(groups, b)
		if sum[rt] < lines[i].y {
			deleted++
			action := lines[i].action
			if action.changed {
				// rollback
				groups[b] = action.from
			}
		}
	}
}

func root(groups []int, i int) int {
	if groups[i] == i {
		return i
	} else {
		return root(groups, groups[i])
	}
}
