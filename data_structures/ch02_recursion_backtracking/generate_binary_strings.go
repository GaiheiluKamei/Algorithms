package rb

import (
	"fmt"
)

/*
Problem-3: Generate all the strings of n bits. Assume A[0...n-1] is an array of size n.
*/

func printResult(A []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
	}

	fmt.Println()
}

func generateBinaryStrings(n int, A []int, i int) {
	if i == n {
		printResult(A, n)
		return 
	}

	// assign "0" at its position and try for all other permutations for remaining positions
	A[i] = 0
	generateBinaryStrings(n, A, i+1)

	// assign "1" at its position and try for all other permutations for remaining positions
	A[i] = 1
	generateBinaryStrings(n, A, i+1)
}

/*
func main() {
	n := 4
	A := make([]int, n)

	generateBinaryStrings(n, A, 0)
}
*/