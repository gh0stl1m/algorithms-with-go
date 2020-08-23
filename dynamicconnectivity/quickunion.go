package dynamicconnectivity

// QuickUnion is the base structure that holds
// the ID slice
type QuickUnion struct {
	ID []uint64
}

// Init method that initialize the dc slice
func (qu *QuickUnion) Init(n uint64) {
	qu.ID = make([]uint64, n)
	for i := range qu.ID {
		qu.ID[i] = uint64(i)
	}
}

// Connected check if two points are linked
func (qu *QuickUnion) Connected(p uint64, q uint64) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}

// Union link two points
func (qu *QuickUnion) Union(p uint64, q uint64) {
	proot := qu.getRoot(p)
	qroot := qu.getRoot(q)

	qu.ID[proot] = qroot
}

func (qu *QuickUnion) getRoot(i uint64) uint64 {
	for i != qu.ID[i] {
		i = qu.ID[i]
	}

	return i
}
