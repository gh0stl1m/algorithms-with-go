package dynamicconnectivity

// QuickFind is the base structure that holds
// the ID slice
type QuickFind struct {
	ID []uint64
}

// Init method that initialize the dc slice
func (qf *QuickFind) Init(n uint64) {
	qf.ID = make([]uint64, n)
	for i := range qf.ID {
		qf.ID[i] = uint64(i)
	}
}

// Connected check if two points are linked
func (qf *QuickFind) Connected(p uint64, q uint64) bool {
	return qf.ID[p] == qf.ID[q]
}

// Union link two points
func (qf *QuickFind) Union(p uint64, q uint64) {
	pid := qf.ID[p]
	qid := qf.ID[q]

	for i := range qf.ID {
		if qf.ID[i] == pid {
			qf.ID[i] = qid
		}
	}
}
