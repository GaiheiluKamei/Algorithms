package rb

import "testing"

func TestHanoi(t *testing.T) {
	Hanoi(3)
	t.Log("Success")
}

/*
func main() {
	hanoi(3)
}

Move disk  1  from peg  A  to peg  C
Move disk  2  from peg  A  to peg  B
Move disk  1  from peg  C  to peg  B
Move disk  3  from peg  A  to peg  C
Move disk  1  from peg  B  to peg  A
Move disk  2  from peg  B  to peg  C
Move disk  1  from peg  A  to peg  C
*/