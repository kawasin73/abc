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

	index := make(map[int]struct{})
	var sum int
	for i := 1; i < 1000; i++ {
		sum += i
		index[sum] = struct{}{}
	}

	fmt.Println(sum)

	for x := 1; x < 499500; x++ {
		if _, ok := index[a + x]; !ok {
			continue
		}
		if _, ok := index[b + x]; !ok {
			continue
		}
		fmt.Println(x)
		break
	}
}
