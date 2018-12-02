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

	var r int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if inside(i, j, n) && inside(i+1, j, n) && inside(i, j+1, n) && inside(i+1, j+1, n) {
				r++
			}
		}
	}
	fmt.Println(r)
}

func inside(x, y, n int) bool {
	if float64(x+y) < float64(n)/2 {
		return false
	} else if float64(x+y) > float64(3*n)/2 {
		return false
	} else if float64(x-y) < float64(-n)/2 {
		return false
	} else if float64(x-y) > float64(n)/2 {
		return false
	}
	return true
}
