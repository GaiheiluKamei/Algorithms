package list

import "testing"
//****************覆盖率有点低***************

func build() *Linked {
	// 9-2-5-3-6
	l := &Linked{}
	l.insertBeginning(2)
	l.insertEnd(5)
	l.insert(3, 6)
	l.insertBeginning(9)
	l.insert(4, 3)

	return l
}

func TestSinglyList(t *testing.T) {
	l := build()

	x, err := l.deleteFirst()
	if err != nil || x.(int) != 9 {
		t.Error(err)
	}

	x, err = l.deleteLast()
	if err != nil || x.(int) != 6 {
		t.Error(err)
	}

	x, err = l.delete(2)
	if err != nil || x.(int) != 5 {
		t.Error(err)
	}

	if l.size != 2 {
		t.Errorf("error size")
	}
}