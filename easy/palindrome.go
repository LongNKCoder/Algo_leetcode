package main

import "fmt"

func main() {
	fmt.Println(solution("abc"))     // false
	fmt.Println(solution("abcdcba")) // true
	fmt.Println(solution("aba"))     // true
	fmt.Println(solution("racecar")) // true
	fmt.Println(solution("asss"))    // false
	fmt.Println(solution("assa"))
	fmt.Println(clearSolution("abc"))     // false
	fmt.Println(clearSolution("abcdcba")) // true
	fmt.Println(clearSolution("aba"))     // true
	fmt.Println(clearSolution("racecar")) // true
	fmt.Println(clearSolution("asss"))    // false
	fmt.Println(clearSolution("assa"))    // true
	fmt.Println(clearSolution("abcdba"))  // true
}

func solution(chains string) bool {
	// check len chains odd or even
	lengthChains := len(chains)
	switch lengthChains {
	case 0, 1:
		return false
	case 2:
		if chains[0] == chains[1] {
			return true
		} else {
			return false
		}
	case 3:
		if chains[0] == chains[2] {
			return true
		} else {
			return false
		}
	}
	if lengthChains%2 == 0 {
		// start at the 0 to middle and the last to middle +1
		for i := 0; i <= lengthChains/2; i++ {
			for j := lengthChains - 1; j >= lengthChains/2+1; j-- {
				if chains[i] != chains[j] {
					return false
				}
				i++
			}
			return true
		}
	} else {
		// start at the 0 to middle and the last to middle +2
		for i := 0; i <= lengthChains/2; i++ {
			for j := lengthChains - 1; j >= lengthChains/2+2; j-- {
				if chains[i] != chains[j] {
					return false
				}
				i++
			}
			return true
		}
	}
	return true
}

func clearSolution(chains string) bool {
	left := 0
	right := len(chains) - 1
	checkSum := 0

	for left < right {
		if chains[left] != chains[right] {
			checkSum++
			continue
		}
		left++
		right--
	}

	if checkSum > 2 {
		return false
	}
	return true
}
