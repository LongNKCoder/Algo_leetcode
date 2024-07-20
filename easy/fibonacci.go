package main

import "fmt"

func main() {
	n := 10
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

//isExist := make(map[int]int, 0)
//fmt.Println(fibonacciCache(n, isExist))
//func fibonacciCache(n int, cache map[int]int) int {
//	if num, ok := cache[n]; ok {
//		return num
//	}
//	if n < 2 {
//		cache[n] = n
//		return n
//	}
//	cache[n] = fibonacciCache(n-1, cache) + fibonacciCache(n-2, cache)
//	return cache[n]
//}
