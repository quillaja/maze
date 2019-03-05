package maze

// NodeSlice is a slice of Nodes.
type NodeSlice []Node

// ToNodeSlice is a convenience method to convert a slice of any type that
// implements Node to a NodeSlice. If slice cannot be converted to []interface{},
// nil is returned.
// func ToNodeSlice(slice interface{}) NodeSlice {
// 	list, ok := slice.([]interface{}) // TODO: something wrong here
// 	if !ok {
// 		return nil
// 	}

// 	s := make(NodeSlice, len(list))
// 	for i := range list {
// 		s[i] = list[i].(Node) // TODO: check assertion and return err to prevent panic?
// 	}
// 	return s
// }

// Append adds the nodes to the slice.
// Must be used as:
//     s = s.Append(n)
func (slice NodeSlice) Append(n ...Node) NodeSlice {
	return append(slice, n...)
}

// AppendUnique adds the node to the slice only if it is not already
// in the slice.
// Must be used as:
//     s = s.AppendUnique(n)
func (slice NodeSlice) AppendUnique(n Node) NodeSlice {
	if !slice.Has(n) {
		slice = slice.Append(n)
	}
	return slice
}

// notFound is used to indicate that a Node was not found in a NodeSlice.
const notFound = -1

// Index returns the first index where 'n' is found, or 'notFound' (-1)
// if the node is not found.
func (slice NodeSlice) index(n Node) int {
	// sort.Search() doesn't work as a general purpose search.
	// return sort.Search(len(slice), func(j int) bool {
	// 	return slice[j] == n
	// })
	for i := 0; i < len(slice); i++ {
		if slice[i] == n {
			return i
		}
	}
	return notFound
}

// Has returns true if the slice contains 'n'.
func (slice NodeSlice) Has(n Node) bool {
	return slice.index(n) != notFound
}

// RemoveAt removes the node at i. The original order of the slice
// is not preserved.
func (slice NodeSlice) removeAt(i int) NodeSlice {
	l := len(slice)
	if l == 0 || i < 0 || i > l-1 { // invalid cases
		return slice
	}

	slice[i] = slice[l-1] // overwrite i with end
	slice[l-1] = nil      // prevent memory leak
	slice = slice[:l-1]   // slice off end

	return slice
}

// Remove removes the first occurence of 'n' from the slice. The
// original order is not preserved.
// Must be used as:
//     s = s.Remove(n)
func (slice NodeSlice) Remove(n Node) NodeSlice {
	i := slice.index(n)
	if i != notFound {
		return slice.removeAt(i)
	}

	return slice
}

// Pop removes and returns the last node on the slice.
func (slice NodeSlice) pop() (NodeSlice, Node) {
	l := len(slice)
	if l == 0 {
		return slice, nil
	}

	n := slice[l-1]
	return slice.removeAt(l - 1), n
}
