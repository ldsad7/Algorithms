package array

import (
	"testing"
)

func TestLen(t *testing.T) {
	h := New(1)
	for i := 0; i < 5; i++ {
		h.Push(Element(i * i))
	}
	if h.Len() != 5 {
		t.Error("Error")
	}
}

func TestGet(t *testing.T) {
	h := New(0)
	for i := 0; i < 10; i++ {
		h.Push(Element(i * i))
	}

	v, err := h.Get(0)
	if v != 0 || err != nil {
		t.Fatal("Get returned an invalid result")
	}

	v, err = h.Get(5)
	if v != 25 || err != nil {
		t.Fatal("Get returned an invalid result")
	}

	v, err = h.Get(9)
	if v != 81 || err != nil {
		t.Fatal("Get returned an invalid result")
	}

	v, err = h.Get(10)
	if err == nil {
		t.Fatal("Get returned an invalid result")
	}

	v, err = h.Get(-1)
	if err == nil {
		t.Fatal("Get returned an invalid result")
	}

	v, err = h.Get(115)
	if err == nil {
		t.Fatal("Get returned an invalid result")
	}
}

func TestSet(t *testing.T) {
	h := New(3)
	h.Push(1)
	err := h.Set(0, 42)
	if err != nil {
		t.Error("Failed to set an array element by valid index")
	}

	err = h.Set(0, 21)
	if err != nil {
		t.Error("Failed to set an array element by valid index")
	}

	err = h.Set(1, 21)
	if err == nil {
		t.Error("Error")
	}

	h.Push(1)
	err = h.Set(1, 21)
	if err != nil {
		t.Error("Error")
	}

	h.Push(1)
	err = h.Set(2, 21)
	if err != nil {
		t.Error("Error")
	}

	h.Push(1)
	h.Push(1)
	err = h.Set(4, 21)
	if err != nil {
		t.Error("Error")
	}
}

func TestInsert(t *testing.T) {
	h := New(3)

	err := h.Insert(0, 1)
	if err != nil || h.elems[0] != 1 {
		t.Error("Error")
	}

	err = h.Insert(0, 2)
	if err != nil || h.elems[0] != 2 {
		t.Error("Error")
	}

	err = h.Insert(5, 3)
	if err == nil {
		t.Error("Error")
	}

	err = h.Insert(3, 4)
	if err == nil {
		t.Error("Error")
	}

	err = h.Insert(2, 4)
	if err != nil {
		t.Error("Error")
	}
}

func TestPush(t *testing.T) {
	h := New(0)

	h.Push(213423)
	v, err := h.Get(h.Len() - 1)
	if err != nil || v != 213423 {
		t.Error("Push test failed")
	}

	h.Push(123411)

	v, err = h.Get(h.Len() - 1)

	if err != nil || v != 123411 {
		t.Error("Push test failed")
	}

	h.Push(123411123)

	v, err = h.Get(2)

	if err != nil || v != 123411123 {
		t.Error("Push test failed")
	}

	v, err = h.Get(1)

	if err != nil || v != 123411 {
		t.Error("Push test failed")
	}
}

func TestDelete(t *testing.T) {
	h := New(0)
	for i := 0; i < 5; i++ {
		h.Push(Element(i * i))
	}

	err := h.Delete(0)

	if err != nil || h.elems[0] != 1 {
		t.Error("Delete Error")
	}

	err = h.Delete(2)

	if err != nil || h.elems[2] != 16 {
		t.Error("Delete Error")
	}

	err = h.Delete(3)

	if err == nil || h.elems[2] != 16 {
		t.Error("Delete Error")
	}

	err = h.Delete(-1)

	if err == nil || h.elems[0] != 1 {
		t.Error("Delete Error")
	}
	h.Delete(0)
	h.Delete(0)
	h.Delete(0)
	h.Delete(0)
}

func TestPop(t *testing.T) {
	h := New(0)
	for i := 0; i < 5; i++ {
		h.Push(Element(i * i))
	}

	err := h.Pop()
	if err != nil || h.elems[3] != 9 {
		t.Error("Pop Error")
	}

	h.Pop()
	h.Pop()
	h.Pop()
	h.Pop()
	err = h.Pop()
	if err == nil {
		t.Error("Pop Error")
	}
}

func TestString(t *testing.T) {
	h := New(0)
	if h.String() != "[]" {
		t.Fatal("String() is invalid for an empty array, expected []")
	}

	h.Push(1)
	if h.String() != "[1]" {
		t.Fatal("String() is invalid for an empty array, expected []")
	}

	h.Delete(0)
	for i := 0; i < 10; i++ {
		h.Push(Element(i * i))
	}

	result := h.String()
	expected := "[0, 1, 4, 9, 16, 25, 36, 49, 64, 81]"

	if result != expected {
		t.Fatalf("String() returned an invalid array representation\nResult: %s\nExpected: %s", result, expected)
	}
}
