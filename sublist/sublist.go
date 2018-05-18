package sublist

import (
	"reflect"
)

// Relation 1
type Relation string

// Sublist 1
func Sublist(listA, listB []int) Relation {
	// first check if a is sublist of B
	if len(listA) < len(listB) && checkSublist(listA, listB) {
		return "sublist"
	}
	if len(listB) < len(listA) && checkSublist(listB, listA) {
		return "superlist"
	}
	if len(listB) == len(listA) && reflect.DeepEqual(listA, listB) {
		return "equal"
	}
	return "unequal"
}

func checkSublist(lista, listb []int) bool {
	size := len(lista)
	for i := 0; i <= len(listb)-size; i++ {
		if reflect.DeepEqual(listb[i:size+i], lista) {
			return true
		}
	}
	return false
}
