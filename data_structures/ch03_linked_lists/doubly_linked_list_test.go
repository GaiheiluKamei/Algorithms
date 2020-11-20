package list

import "testing"

func buildDL() *dl {
	dl := &dl{}

	dl.insertBeginning(100)
	dl.insertBeginning(90)
	dl.insertBeginning(70)
	dl.insertBeginning(60)
	dl.insertBeginning(60)
	dl.insertBeginning(50)
	dl.insertBeginning(40)
	dl.insertBeginning(30)
	dl.insertBeginning(19)

	return dl
}

func TestDoubleList(t *testing.T) {
	dl := buildDL()
	dl.traverse()

	dl.deleteLast()
	dl.traverse()

	dl.delete(3)
	dl.traverse()

	dl.deleteFirst()
	dl.traverse()

	t.Log("Success")
}
