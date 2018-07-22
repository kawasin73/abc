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

	var min int
	var max int
	for i := 0; i < n; i++ {
		a := nextInt()
		if i == 0 {
			min = a
			max = a
			continue
		}
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	fmt.Println(max - min)
}
