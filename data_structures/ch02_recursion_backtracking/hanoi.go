package rb

import "fmt"
/*
Problem-1: Discuss Towers of Hanoipuzzle.

Solution: 
The Towers of Hanoi is a mathematical puzzle. It consists of three rods(or pegs or towers), and a number of disks of different sizes which can slide onto any rod. The puzzle starts with the disks on one rod in ascending order of size, the smallest at the top, thus making a conical shape.The objective of the puzzle is to move the entire stack to another rod, satisfyingthe following rules:

- Only one disk may be moved at a time.
- Each move consists of taking the upper disk from one of the rods and sliding it onto another rod, on top of the other disks that may already be present on that rod.
- No disk may be placed on top of a smaller disk.

Algorithm:
- Move the top n݊−1 disks from ݁ܿ`Source` to `Auxiliary` tower
- Move the nth disk from the `Source` to `Destination` tower
- Move the n-1 disks from `Auxiliary` tower to `Destination` tower
- Transferring the top n-1 disks from `Source` to `Auxiliary` tower can again be thought of as a fresh problem and can be solved in the same manner. Once we solve Towers of Hanoi with three disks, we can solve it with any number of disks with the above algorithm.
*/

func hanoi(n int, from, to, tmp string) {
	// if only 1 disk, make the move and return
	if n == 1 {
		fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)
		return
	}

	// step 1: move top n-1 disks from A to B, using C as auxiliary
	hanoi(n-1, from, tmp, to)

	// step 2: move remaining disks from A to C
	fmt.Println("Move disk ", n, " from peg ", from, " to peg ", to)

	// step 3: move n-1 disks from B to C, using A as auxiliary
	hanoi(n-1, tmp, to, from)
}

// Hanoi solve problem of hanoi.
func Hanoi(n int) {
	hanoi(n, "A", "C", "B")
}