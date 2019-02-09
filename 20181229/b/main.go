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

	l := nextInt()
	n := nextInt()

	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}

	head := 0
	tail := n - 1
	cur := 0

	

	var length int
	for i := 0; i < n; i++ {
		headlength := (x[head] - cur + l) % l
		taillength := (cur - x[tail] + l) % l

		//fmt.Println("len", headlength, taillength)
		if headlength > taillength {
			fmt.Println("head", headlength)
			length += headlength
			cur = x[head]
			head++
		} else {
			fmt.Println("tail", taillength)
			length += taillength
			cur = x[tail]
			tail--
		}
	}

	fmt.Println(length)
}
