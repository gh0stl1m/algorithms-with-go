package hashmap

// Entity reprensents the element inside the hashmap
type Entity struct {
	Key   uint
	Value interface{}
	Next  *Entity
}

// HashTable reprensents the elements inside the hash map
type HashTable struct {
	ArrayHash []Entity
	Size      uint
}

// Init initialize the ht
func (ht *HashTable) Init(size uint) *HashTable {
	return &HashTable{
		ArrayHash: make([]Entity, size),
		Size:      size,
	}
}

// Put adds elements to the hash table
func (ht *HashTable) Put(key uint, value interface{}) {
	var hashIndex uint = ht.getHash(key)
	var arrayValue Entity = ht.ArrayHash[hashIndex]
	var newItem Entity = Entity{
		Key:   key,
		Value: value,
		Next:  arrayValue.Next,
	}

	arrayValue.Next = &newItem
}

// Get reads an element from the has table
func (ht *HashTable) Get(key uint) interface{} {
	var value interface{}
	var hashIndex uint = ht.getHash(key)
	var arrayValue Entity = ht.ArrayHash[hashIndex]

	for arrayValue.Value != nil {
		if arrayValue.Key == key {
			value = arrayValue.Value
			break
		}

		arrayValue = *arrayValue.Next
	}

	return value
}

func (ht *HashTable) getHash(key uint) uint {
	return key % ht.Size
}

// FindUniqueNumber search the unique value in an array
func FindUniqueNumber(array []int) int {
	var numberCounter map[int]int = make(map[int]int, len(array))
	var uniqueNumber int

	for _, value := range array {
		numberCounter[value]++
	}

	for key, value := range numberCounter {
		if value == 1 {
			uniqueNumber = key
		}
	}

	return uniqueNumber
}
