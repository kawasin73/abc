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

var sosuu map[int]struct{} = make(map[int]struct{})

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	pool := make([]int, 101)
	for i := 2; i <= n; i++ {
		ok := true
		for k, _ := range sosuu {
			i2 := i
			var l int
			for i2%k == 0 && i2 > 0 {
				l ++
				i2 /= k
			}
			if l > 0 {
				ok = false
				pool[k] += l
			}
		}
		if ok {
			sosuu[i] = struct{}{}
			pool[i] = 1
		}
	}

	var more3 int
	for i := 2; i <= n; i++ {
		if pool[i] >= 2 {
			more3++
		}
	}

	var more5 int
	for i := 2; i <= n; i++ {
		if pool[i] >= 4 {
			more5++
		}
	}

	var more15 int
	for i := 2; i <= n; i++ {
		if pool[i] >= 14 {
			more15++
		}
	}

	var more25 int
	for i := 2; i <= n; i++ {
		if pool[i] >= 24 {
			more25++
		}
	}

	var more75 int
	for i := 2; i <= n; i++ {
		if pool[i] >= 74 {
			more75++
		}
	}

	//fmt.Println(pool[0:n])
	//fmt.Println("more75", more75)
	//fmt.Println("more25", more25)
	//fmt.Println("more15", more15)
	//fmt.Println("more5", more5)
	//fmt.Println("more3", more3)

	result := 0
	if more5 >= 3 {
		result += more5 * (more5 - 1) * (more5 - 2) / 2
	}
	if more5 >= 2 && more3 >= 3 {
		result += (more3 - more5) * more5 * (more5 - 1) / 2
	}
	if more15 >= 1 && more5 >= 2 {
		result += more15*(more15-1) + more15*(more5-more15)
	}
	if more25 >= 1 && more3 >= 2 {
		result += more25*(more25-1) + more25*(more3-more25)
	}
	if more75 >= 1 {
		result += more75
	}

	fmt.Println(result)
}
