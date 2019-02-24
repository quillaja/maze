package maze

// NodeSlice is a slice of Nodes.
type NodeSlice []Node

// ToNodeSlice is a convenience method to convert a slice of any type that
// implements Node to a NodeSlice.
func ToNodeSlice(list []interface{}) NodeSlice {
	s := make(NodeSlice, len(list))
	for i := range list {
		s[i] = list[i].(Node) // TODO: check assertion and return err to prevent panic?
	}
	return s
}

// Append adds the nodes to the slice.
func (slice NodeSlice) Append(n ...Node) NodeSlice {
	return append(slice, n...)
}

// AppendUnique adds the node to the slice only if it is not already
// in the slice.
func (slice NodeSlice) AppendUnique(n Node) NodeSlice {
	if !slice.Has(n) {
		slice = slice.Append(n)
	}
	return slice
}

// Index returns the first index where 'n' is found, or 'len(slice)'
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
	return len(slice)
}

// Has returns true if the slice contains 'n'.
func (slice NodeSlice) Has(n Node) bool {
	return slice.index(n) != len(slice)
}

// RemoveAt removes the node at i. The original order of the slice
// is not preserved.
func (slice NodeSlice) removeAt(i int) NodeSlice {
	l := len(slice)
	if l == 0 {
		return slice
	}

	slice[i] = slice[l-1] // overwrite i with end
	slice[l-1] = nil      // prevent memory leak
	slice = slice[:l-1]   // slice off end

	return slice
}

// Remove removes the first occurence of 'n' from the slice. The
// original order is not preserved.
func (slice NodeSlice) Remove(n Node) NodeSlice {
	i := slice.index(n)
	if i < len(slice) {
		return slice.removeAt(i)
	}

	return slice
}

// Pop removes and returns the last node on the slice.
func (slice NodeSlice) pop() (NodeSlice, Node) {
	n := slice[len(slice)-1]
	return slice.removeAt(len(slice) - 1), n
}
