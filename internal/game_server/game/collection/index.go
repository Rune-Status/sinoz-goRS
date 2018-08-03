package collection

import (
	"errors"
)

type Index struct {
	elements []Element
	offset   int
	size     int
	limit    int
}

type Element interface{}

func NewIndex(offset, capacity int) *Index {
	return &Index{elements: make([]Element, capacity+offset), offset: offset}
}

func (idx *Index) Set(slot int, element Element) error {
	if slot < idx.offset || slot >= len(idx.elements) {
		return errors.New("index out of bounds")
	}

	if element != nil {
		if slot >= idx.limit {
			idx.limit += 1
		}

		if idx.elements[slot] == nil {
			idx.size += 1
		}
	} else {
		if slot >= idx.limit {
			idx.limit -= 1
		}

		if idx.elements[slot] != nil {
			idx.size -= 1
		}
	}

	idx.elements[slot] = element
	return nil
}

func (idx *Index) IndexWhere(p func(Element) bool) int {
	for i := idx.offset; i <= idx.limit; i++ {
		element := idx.elements[i]
		if p(element) {
			return i
		}
	}

	return -1
}

func (idx *Index) Count(p func(Element) bool) int {
	var result int

	for i := idx.offset; i <= idx.limit; i++ {
		element := idx.elements[i]
		if element != nil && p(element) {
			result += 1
		}
	}

	return result
}

func (idx *Index) Find(p func(Element) bool) Element {
	for i := idx.offset; i <= idx.limit; i++ {
		element := idx.elements[i]
		if element != nil && p(element) {
			return element
		}
	}

	return nil
}

func (idx *Index) Get(slot int) (Element, error) {
	if slot < idx.offset || slot >= len(idx.elements) {
		return nil, errors.New("index out of bounds")
	}

	return idx.elements[slot], nil
}

func (idx *Index) forEach(f func(Element)) {
	for i := idx.offset; i <= idx.limit; i++ {
		element := idx.elements[i]
		if element != nil {
			f(element)
		}
	}
}

func (idx *Index) GetSize() int {
	return idx.size
}
