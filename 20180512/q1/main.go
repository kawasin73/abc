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
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()

	d2 := a - c
	if d2 < 0 {
		d2 = -d2
	}

	if d2 <= d {
		fmt.Println("Yes")
		return
	}

	d2 = a - b
	if d2 < 0 {
		d2 = -d2
	}

	if d2 > d {
		fmt.Println("No")
		return
	}

	d2 = b - c
	if d2 < 0 {
		d2 = -d2
	}

	if d2 > d {
		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
