package fulltext

import (
	"sort"
	"strings"
)

// Index implements fulltext search
type Index struct {
	m map[uint8]*Index
	set[int]
}

// New creates a fulltext search index for the given documents
func New(docs []string) *Index {
	m := make(map[string]map[int]bool)

	for i := 0; i < len(docs); i++ {
		words := strings.Fields(docs[i])
		for j := 0; j < len(words); j++ {
			if m[words[j]] == nil {
				m[words[j]] = make(map[int]bool)
			}
			m[words[j]][i+1] = true
		}
	}
	return &(Index{m})
}

// Search returns a slice of unique ids of documents that contain all words from the query.
func (idx *Index) Search(query string) []int {
	if query == "" {
		return []int{}
	}
	ret := make(map[int]bool)
	arr := strings.Fields(query)
	fl := 0
	for i := range arr {
		if idx.m[arr[i]] == nil {
			return []int{}
		}
		if fl == 0 {
			for value := range idx.m[arr[i]] {
				ret[value] = true
			}
			fl = 1
		} else {
			tmp := make(map[int]bool)
			for value := range ret {
				if idx.m[arr[i]][value] == true {
					tmp[value] = true
				}
			}
			ret = tmp
		}
	}
	ret_arr := []int{}
	for value := range ret {
		ret_arr = append(ret_arr, value-1)
	}
	sort.Ints(ret_arr)
	return ret_arr
}
