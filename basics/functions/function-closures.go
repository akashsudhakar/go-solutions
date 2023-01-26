package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	first := 0
	second := 1
	return func() int {
		first, second = second, first+second
		return second
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	f := fibonacci()
	i := 0
	if i < 2 {
		fmt.Println(i)
		i += 1
	}
	for i = 2; i < 10; i++ {
		fmt.Println(f())
	}
}
