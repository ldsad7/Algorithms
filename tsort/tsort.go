package tsort

import (
	"fmt"
)

// THE FILE MUST CONTAIN A SINGLE FUNCTION DEFINITION

// TSort performs a topological sort on the given graph represented as a list of edges
func TSort(edges []string) ([]string, error) {
	if len(edges)%2 == 1 {
		return nil, fmt.Errorf("Error: odd number of elements")
	}
	L := []string{}
	nodes := make(map[string]bool)
	m := make(map[string]map[string]bool)
	for i := 0; 2*i < len(edges); i++ {
		nodes[edges[2*i]] = true
		nodes[edges[2*i+1]] = true
		if m[edges[2*i]] == nil {
			m[edges[2*i]] = make(map[string]bool)
		}
		m[edges[2*i]][edges[2*i+1]] = true
	}
	for len(nodes) > 0 {
		fl := 0
		for value := range nodes {
			tmp := 0
			for node := range m[value] {
				if nodes[node] == true {
					tmp = 1
					break
				}
			}
			if tmp == 0 {
				L = append(L, value)
				delete(nodes, value)
				fl = 1
			}
		}
		if fl == 0 && len(nodes) > 0 {
			return nil, fmt.Errorf("Error: cyclic graph")
		}
	}
	for i := 0; i < len(L)/2; i++ {
		j := len(L) - i - 1
		L[i], L[j] = L[j], L[i]
	}
	return L, nil
}
