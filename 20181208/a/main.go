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
	d := nextInt()
	var v string
	switch d {
	case 22:
		v = "Christmas Eve Eve Eve"
	case 23:
		v = "Christmas Eve Eve"
	case 24:
		v = "Christmas Eve"
	case 25:
		v = "Christmas"
	}
	fmt.Println(v)
}
