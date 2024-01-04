package avl

import (
	"testing"
)

func TestSetNewTreeSet(t *testing.T) {
	var set = NewTreeSet[int]()
	if set == nil {
		panic("Bad NewTreeSet")
	}
}

func TestSetPutRemove(t *testing.T) {
	var set = NewTreeSet[int]()
	var i = 0
	set.Put(&i)
	if set.root == nil || set.Size() != 1 {
		panic("Bad Put Remove")
	}
	set.Remove(&i)
	if set.root != nil || set.Size() != 0 {
		panic("Bad Put Remove")
	}
}

func TestSetPutFind(t *testing.T) {
	var set = NewTreeSet[int]()
	var i = 0
	set.Put(&i)
	var r = set.Find(&i)
	if i != *r {
		panic("Bad Put Find")
	}
}

func TestSetPutFindMin(t *testing.T) {
	var set = NewTreeSet[int]()
	var r = set.FindMin()
	if r != nil {
		panic("Bad Put FindMin")
	}
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		set.Put(&a[i])
	}
	r = set.FindMin()
	if a[0] != *r || set.Size() != 8 {
		panic("Bad Put FindMin")
	}
}

func TestSetPutRemoveMin(t *testing.T) {
	var set = NewTreeSet[int]()
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		set.Put(&a[i])
	}
	set.RemoveMin()
	var r = set.Find(&a[0])
	if r != nil || set.Size() != 7 {
		panic("Bad Put RemoveMin")
	}
}

func TestSetPutFindMax(t *testing.T) {
	var set = NewTreeSet[int]()
	var r = set.FindMax()
	if r != nil {
		panic("Bad Put FindMax")
	}
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		set.Put(&a[i])
	}
	r = set.FindMax()
	if a[len(a)-1] != *r || set.Size() != 8 {
		panic("Bad Put FindMax")
	}
}

func TestSetPutRemoveMax(t *testing.T) {
	var set = NewTreeSet[int]()
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		set.Put(&a[i])
	}
	set.RemoveMax()
	var r = set.Find(&a[len(a)-1])
	if r != nil || set.Size() != 7 {
		panic("Bad Put RemoveMax")
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
