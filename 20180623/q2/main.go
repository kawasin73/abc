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
	s := nextString()
	n, _ := strconv.Atoi(s)
	i := n
	var sn int
	for i > 0 {
		sn += i % 10
		i = i / 10
	}
	if n%sn == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
