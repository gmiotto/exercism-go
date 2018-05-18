package tree

import (
	"errors"
	"sort"
)

// Record 1
type Record struct {
	ID, Parent int
}

// Node 1
type Node struct {
	ID       int
	Children []*Node
}

// Build 1
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records[:], func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	mapRecordNode := make(map[int]*Node)
	var root *Node

	for i := 0; i < len(records); i++ {
		if records[i].ID != i {
			return nil, errors.New("Records ID are no continuous")
		}
		if i == 0 {
			if records[i].ID != 0 || records[i].Parent != 0 {
				return nil, errors.New("No root node")
			}
			root = &Node{}
			mapRecordNode[i] = root
			continue
		}
		if records[i].ID < records[i].Parent {
			return nil, errors.New("Parent smaller than itself ID")
		}
		if val, ok := mapRecordNode[records[i].Parent]; ok {
			if _, ok1 := mapRecordNode[records[i].ID]; ok1 {
				return nil, errors.New("Duplicate node is not allowed")
			}
			newNode := &Node{records[i].ID, nil}
			val.Children = append(val.Children, newNode)
			mapRecordNode[records[i].ID] = newNode
		} else {
			return nil, errors.New("Record " + string(records[i].ID) + "has an invalid parent node")
		}
	}

	return root, nil
}
