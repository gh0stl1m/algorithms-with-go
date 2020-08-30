package dynamicconnectivity

// WeightedQuickUnion is the base structure that holds
// the ID slice
type WeightedQuickUnion struct {
	ID   []uint64
	Size []uint64
}

// Init method that initialize the dc slice
func (wqu *WeightedQuickUnion) Init(n uint64) {
	wqu.ID = make([]uint64, n)
	for i := range wqu.ID {
		wqu.ID[i] = uint64(i)
		wqu.Size[i] = uint64(i)
	}
}

// Connected check if two points are linked
func (wqu *WeightedQuickUnion) Connected(p uint64, q uint64) bool {
	return wqu.getRoot(p) == wqu.getRoot(q)
}

// Union link two points
func (wqu *WeightedQuickUnion) Union(p uint64, q uint64) {
	proot := wqu.getRoot(p)
	qroot := wqu.getRoot(q)

	if proot == qroot {
		return
	}

	if wqu.Size[proot] < wqu.Size[qroot] {
		wqu.ID[proot] = qroot
		wqu.Size[qroot] += wqu.Size[proot]
	} else {
		wqu.ID[qroot] = proot
		wqu.Size[proot] += wqu.Size[qroot]
	}

}

func (wqu *WeightedQuickUnion) getRoot(i uint64) uint64 {
	for i != wqu.ID[i] {
		i = wqu.ID[i]
	}

	return i
}
