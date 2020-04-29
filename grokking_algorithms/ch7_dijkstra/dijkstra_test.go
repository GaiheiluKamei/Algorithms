package dijkstra

import (
	"reflect"
	"testing"
)

func TestFindLowestCostNode(t *testing.T) {
	costs := map[string]int {
		"A": 6,
		"B": 2,
		"FIN": 100000,
	}

	processed := make(map[string]struct{})

	if res := findLowestCostNode(costs, processed); res != "B" {
		t.Errorf("expect 'B' but got %v", res)
	}
}

func TestDijkstra(t *testing.T) {
	graph := map[string]map[string]int{
		"START": {"A": 6, "B": 2},
		"A": {"C": 1},
		"B": {"A": 3, "C": 5, "FIN": 6},
		"C": {"FIN": 1},
		"FIN": {},
	}

	costs := map[string]int{
		"A": 6,
		"B": 2,
		"C": 10000,
		"FIN": 10000,
	}

	parents := map[string]string{
		"A": "START",
		"B": "START",
		"FIN": "",
	}

	c, p := Dijkstra(graph, costs, parents)
	// fmt.Println(c, p)
	if !reflect.DeepEqual(c, costs) || !reflect.DeepEqual(p, parents) {
		t.Errorf("got error costs, parents: %v, %v", c, p)
	}
}
