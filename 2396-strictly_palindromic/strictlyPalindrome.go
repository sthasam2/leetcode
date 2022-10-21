package strictlyPalindromic

import (
	"fmt"
)

///////////////////////
// SOLUTION 1
///////////////////////

// func DecimalToBaseAndReverse(n int, b int) (string, string) {
// 	var conv, rev string = "", ""
// 	quo := n
// 	for quo != 0 {
// 		rem := strconv.Itoa(quo % b)
// 		quo = quo / b
// 		conv, rev = conv+rem, rem+rev
// 	}
// 	return conv, rev
// }

// func isStrictlyPalindromic(n int) bool {

// 	if n < 4 {
// 		return false
// 	}

// 	for i := 2; i < n-1; i++ {
// 		if c, r := DecimalToBaseAndReverse(n, i); c != r {
// 			return false
// 		}
// 	}
// 	return true
// }

///////////////////////
// SOLUTION 2
///////////////////////

// func isStrictlyPalindromic(n int) bool {
// 	if n < 5 {
// 		return false
// 	}

// 	for i := 2; i < n-1; i++ {
// 		var c, r string = "", ""
// 		q := n
// 		for q != 0 {
// 			rem := strconv.Itoa(q % i)
// 			q = q / i
// 			c, r = c+rem, rem+r
// 		}

// 		if c != r {
// 			return false
// 		}
// 	}

//		return true
//	}

///////////////////////
// SOLUTION 3
///////////////////////

func isStrictlyPalindromic(n int) bool {
	if n < 5 {
		return false
	}

	for i := n - 2; i > 1; i-- {
		q, t, rev := n, 0, 0

		for q > 0 {
			q = q / i
			rev += (i ^ t) * (q % i)
		}

		if n != rev {
			return false
		}
	}

	return true
}

///////////////////////
// DRIVER
///////////////////////

func main() {
	input := 11
	fmt.Println("Input:\t", input)
	fmt.Println("Is Strictly Palindromic?: ", isStrictlyPalindromic(11))
}
