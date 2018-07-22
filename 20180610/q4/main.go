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
	c := nextInt()

	d := make([][]int, c)
	for i := 0; i < c; i++ {
		d[i] = make([]int, c)
		for j := 0; j < c; j++ {
			d[i][j] = nextInt()
		}
	}

	m0 := make(map[int]int)
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			key := (i + j) % 3
			var m map[int]int
			switch key {
			case 0:
				m = m0
			case 1:
				m = m1
			case 2:
				m = m2
			}
			v := nextInt()
			v--
			l, _ := m[v]
			m[v] = l + 1
		}
	}

	if n == 1 {
		fmt.Println(0)
		return
	}

	pair0 := findMin(m0, c, d)
	pair1 := findMin(m1, c, d)
	pair2 := findMin(m2, c, d)

	fmt.Println(len(m0))
	fmt.Println(len(m1))
	fmt.Println(len(m2))

	fmt.Println("pair0", pair0)
	fmt.Println("pair1", pair1)
	fmt.Println("pair2", pair2)

	var min = 0
	for i := 0; i < len(pair0); i++ {
		for j := 0; j < len(pair1); j++ {
			for l := 0; l < len(pair2); l++ {
				if pair0[i].v == pair2[l].v {
					continue
				} else if pair0[i].v == pair1[j].v {
					continue
				} else if pair1[j].v == pair2[l].v {
					continue
				}
				hoge := pair0[i].cost + pair1[j].cost + pair2[l].cost
				if min == 0 || hoge < min {
					min = hoge
				}
			}
		}
	}
	fmt.Println(min)
}

type pair struct {
	v    int
	cost int
}

func findMin(m map[int]int, c int, d [][]int) []pair {
	var max int
	result := make([]pair, 1, 4)
	result[0].v = -1
	for i := 0; i < c; i++ {
		for k, v := range m {
			cost := d[k][i] * v
			if len(result) == 4 {
				if cost < max {
					result[3] = pair{v: i, cost: cost}
					for i := 2; i > 0; i-- {
						if result[i-1].cost > result[i].cost {
							result[i-1], result[i] = result[i], result[i-1]
						}
					}
				}
			} else {
				result = append(result, pair{v: i, cost: cost})
				for i := len(result)-1; i > 0; i-- {
					if result[i-1].cost > result[i].cost {
						result[i-1], result[i] = result[i], result[i-1]
					}
				}
			}
			max = cost
		}
	}
	return result
}
