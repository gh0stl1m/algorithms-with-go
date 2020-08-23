package dynamicconnectivity

// UnionFind defines the behavior of the dynamic
// connectivity
type UnionFind interface {
	Union(p uint64, q uint64)
	Connected(p uint64, q uint64) bool
}
