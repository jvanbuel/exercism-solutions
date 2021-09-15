package tree

import (
	"errors"
)

// Record is a record in the database containing a node ID and a parent ID
type Record struct {
	ID     int
	Parent int
}

// Node is a node of the tree containing an ID and a list of children
type Node struct {
	ID       int
	Children []*Node
}

//Build builds a tree from a list of records
func Build(r []Record) (*Node, error) {
	if len(r) == 0 {
		return nil, nil
	}

	if len(r) == 1 && r[0].ID < r[0].Parent {
		return nil, errors.New("Parent ID should always be smaller than child ID")
	}

	if len(r) == 1 {
		return &Node{
			ID:       r[0].ID,
			Children: []*Node(nil),
		}, nil
	}

	//	find lowest id larger
	lID := lowestID(r)

	// append each record ID where Parent == lowest id to ID list of children
	children := []*Node(nil) // TODO: figure out why nil needs to be cast to type []*Node for deepEqual to work correctly
	for _, v := range r {
		if v.Parent == lID && v.ID != 0 {
			// for each child, build tree, by filtering out smaller IDs
			child, err := Build(filterSmallerIDs(v.ID, r))
			if err != nil {
				return nil, errors.New("something went wrong")
			}
			children = append(children, child)

		}
	}
	return &Node{
		ID:       lID,
		Children: children,
	}, nil
}

func filterSmallerIDs(ID int, r []Record) []Record {
	filtered := []Record{}
	for _, rec := range r {
		if rec.ID >= ID {
			filtered = append(filtered, rec)
		}
	}
	return filtered
}

func lowestID(r []Record) int {
	lID := r[0].ID
	for _, rec := range r {
		if rec.ID < lID {
			lID = rec.ID
		}
	}
	return lID
}

func sort(r []Record)
