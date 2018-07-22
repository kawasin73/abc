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
	k := nextInt()

	i := 0
	j := 0
	list := make([]int, 20)
	list[0] = 1

	for {
		if isOk(list, i, j) {
			printL(list, i)
			k--
			if k <= 0 {
				break
			}
		} else {
			list[j] = 9
			if j < i {
				j++
				continue
			}
		}
		if next(list, i, j) {
			i++
			j = 0
		}
	}
}

func isOk(list []int, i, j int) bool {
	n := calcN(list, i)
	sn := calcSn(list, i)
	for m := 0; m < j; m++ {
		sn *= 10
	}
	return n <= sn
}

func next(list []int, i int, j int) (top bool) {
	if list[j] == 9 {
		if i == j {
			for k := 0; k <= i; k++ {
				list[k] = 0
			}
			list[i+1] = 1
			return true
		} else {
			for k := j; true; k++ {
				if list[k] == 9 {
					list[k] = 0
				} else {
					list[k]++
					break
				}
			}
			if list[i+1] == 1 {
				return true
			}
			return false
		}
	} else {
		list[j]++
	}
	return false
}

func printL(list []int, i int) {
	for m := i; m >= 0; m-- {
		fmt.Print(list[m])
	}
	fmt.Println()
}

func calcN(list []int, i int) int {
	var sum int
	for k := i; k >= 0; k-- {
		sum *= 10
		sum += list[k]
	}
	return sum
}

func calcSn(list []int, i int) int {
	var sum int
	for _, v := range list[:i+1] {
		sum += v
	}
	return sum
}
