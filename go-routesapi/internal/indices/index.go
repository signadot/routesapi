package indices

import (
	"sort"
)

type KeySet[Key comparable] map[Key]bool

// An index maps strings to a set of keys. Keys are the canonical
// representation of the finest grain unit we use to track changes.
type Index[Key comparable] map[string]KeySet[Key]

// adds a key k to be indexed by value v
func (i Index[Key]) Add(v string, k Key) {
	dd, ok := i[v]
	if !ok {
		dd = KeySet[Key]{}
		i[v] = dd
	}
	dd[k] = true
}

// get the set of keys indexed by value v
func (i Index[Key]) Get(v string) KeySet[Key] {
	return i[v]
}

// removes a key k from being indexed by value v
func (i Index[Key]) Remove(v string, k Key) {
	dd, ok := i[v]
	if !ok {
		return
	}
	delete(dd, k)
}

// efficient intersection of a slice of sets of keys this is always called with
// a non-empty slice of keys.
func KeySetIntersect[Key comparable](keys ...KeySet[Key]) KeySet[Key] {
	if len(keys) == 0 {
		panic("no keys")
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	res := keys[0]
	for j := 1; j < len(keys); j++ {
		for k := range res {
			if !keys[j][k] {
				delete(res, k)
			}
		}
	}
	return res
}
