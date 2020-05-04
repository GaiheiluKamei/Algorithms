package offer

import (
	"reflect"
	"testing"
)

func TestReplaceBlank(t *testing.T) {
	str := []byte{'a', ' ', 'b', 'c', ' ', 'X', 'X', 'X', 'X'}  // 'X'作为占位符，保证后面有足够空间容纳增加的字符
	ReplaceBlank(str, 5)
	if reflect.DeepEqual(string(str), []byte{'a', '%', '2', '0', 'b', 'c', '%', '2', '0'}) {
		t.Error("I will eat shit online if this print out.")
	}
}
