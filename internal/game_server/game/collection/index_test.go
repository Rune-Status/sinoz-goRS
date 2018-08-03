package collection

import (
	"testing"
	"strconv"
)

func TestIndex_Set(t *testing.T) {
	index := NewIndex(1, 16)
	index.Set(1, "hello world")

	value, _ := index.Get(1)
	if value.(string) != "hello world" {
		t.Fatalf("extracted value at slot 1 did not equal %v but actually equals %v", "hello world", value)
	}

	size := index.GetSize()
	if size != 1 {
		t.Fatalf("size expected to be 1, is actually %v", size)
	}
}

func TestIndex_Find(t *testing.T) {
	index := NewIndex(1, 16)

	for i := 1; i < 10; i++ {
		index.Set(i, "hello world " + strconv.Itoa(i))
	}

	element := index.Find(func(element Element) bool { return element.(string) == "hello world 5" })
	if element != "hello world 5" {
		t.Fatalf("element expected to be 'hello world 5' but was %v instead", element)
	}
}

func TestIndex_Count(t *testing.T) {
	index := NewIndex(1, 16)

	for i := 1; i < 10; i++ {
		index.Set(i, "hello world")
	}

	count := index.Count(func(element Element) bool { return element.(string) == "hello world" })
	if count != 9 {
		t.Fatalf("index expected to be 11 but was %v instead", count)
	}
}

func TestIndex_IndexWhere(t *testing.T) {
	index := NewIndex(1, 16)
	index.Set(1, "hello world")

	expectedSlot := index.IndexWhere(func(element Element) bool { return element.(string) == "hello world" })
	if expectedSlot != 1 {
		t.Fatalf("index expected to be 2 but was %v instead", expectedSlot)
	}
}