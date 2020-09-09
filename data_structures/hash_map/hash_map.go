package hashmap

import (
	"strings"
)

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

// Problem: Given an array with numbers, find the unique one

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

// Problem: Key phrase problem TF-IDF find the most important word
//          excluding the not relevant words

// GetMostRelevantWord returns the most important word given a phrase
func GetMostRelevantWord(phrase string, wordsToExclude []string) string {
	var mostFrequentWord string
	var mostFrequentWordTimes int = 0
	var wordsFrecuency map[string]int = make(map[string]int)
	var textSplitted []string = strings.Split(phrase, " ")

	for _, word := range textSplitted {
		if !contains(word, wordsToExclude) {
			wordsFrecuency[word]++
			if wordsFrecuency[word] > mostFrequentWordTimes {
				mostFrequentWord = word
				mostFrequentWordTimes = wordsFrecuency[word]
			}
		}
	}

	return mostFrequentWord
}

func contains(word string, source []string) bool {
	for _, value := range source {
		if word == value {
			return true
		}
	}
	return false
}
