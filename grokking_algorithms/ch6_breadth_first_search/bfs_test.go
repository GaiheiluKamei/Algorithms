package bfs

import "testing"

func TestAddEdge(t *testing.T) {
	g := &Graph{}
	n1 := &Node{"Jack", false}
	n2 := &Node{"Tim", true}

	g.AddEdge(n1, n2)

	if g.Edges[n1][0] != n2 {
		t.Errorf("expect %v, got %v", n2, g.Edges[n1][0])
	}
}

func TestBFS(t *testing.T) {
	g := &Graph{}
	jojo := &Node{Name: "jojo", MangoSeller: false}
	claire := &Node{Name: "claire", MangoSeller: false}
	ak := &Node{Name: "jonny", MangoSeller: false}
	thom := &Node{Name: "thom", MangoSeller: false}
	alice := &Node{Name: "alice", MangoSeller: false}
	peggy := &Node{Name: "peggy", MangoSeller: false}
	bob := &Node{Name: "bob", MangoSeller: false}
	J20 := &Node{Name: "anuj", MangoSeller: false}

	g.AddEdge(jojo, claire)
	g.AddEdge(claire, thom)
	g.AddEdge(claire, ak)
	g.AddEdge(jojo, alice)
	g.AddEdge(alice, peggy)
	g.AddEdge(jojo, bob)
	g.AddEdge(bob, J20)
	g.AddEdge(bob, peggy)

	t.Run("J20 is the mango seller", func(t *testing.T) {
		 J20.MangoSeller = true

		 n := g.BFS(jojo)

		 if n != J20 {
		 	t.Errorf("expect %v, got %v", J20, n)
		 }

		 J20.MangoSeller = false
	})

	t.Run("ak is the mango seller", func(t *testing.T) {
		ak.MangoSeller = true

		n := g.BFS(jojo)

		if n != ak {
			t.Errorf("expect %v, got %v", ak, n)
		}

		ak.MangoSeller = false
	})
}

func TestBFSNil(t *testing.T) {
	g := &Graph{}

	jojo := &Node{Name: "jojo", MangoSeller: false}
	claire := &Node{Name: "claire", MangoSeller: false}

	g.AddEdge(jojo, claire)
	g.AddEdge(claire, jojo)

	n := g.BFS(jojo)
	if n != nil {
		t.Errorf("expect %v, got %v", nil, n)
	}
}
