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
	a := nextString()
	b := nextString()
	c := nextString()

	var count int
	for i := 0; i < n; i++ {
		if a[i] == b[i] && b[i] == c[i] {
			continue
		} else if a[i] == b[i] || b[i] == c[i] || a[i] == c[i] {
			count++
		} else {
			count += 2
		}
	}

	fmt.Println(count)
}
