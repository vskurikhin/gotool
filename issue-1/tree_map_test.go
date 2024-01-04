package avl

import (
	"testing"
)

func TestTreeMapNewTreeMap(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	if tree == nil {
		panic("Bad NewTreeSet")
	}
}

func TestTreeMapPutRemove(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var i = 0
	tree.Put(&i, &i)
	if tree.root == nil || tree.Size() != 1 {
		panic("Bad Put Remove")
	}
	tree.Remove(&i)
	if tree.root != nil || tree.Size() != 0 {
		panic("Bad Put Remove")
	}
}

func TestTreeMapPutGetValues(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var r = tree.GetValues(nil)
	if r != nil {
		panic("Bad Put GetValues")
	}
	var i = 0
	tree.Put(&i, &i)
	r = tree.GetValues(&i)
	if r == nil || *r[0] != i || tree.Size() != 1 {
		panic("Bad Put GetValues")
	}
}

func TestTreeMapPutValueToMap(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var a = [...]int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7}
	var b = [...]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		tree.Put(&a[i], &b[i])
	}
	if tree.root == nil || tree.Size() != 8 {
		panic("Bad Put GetValues")
	}
	var r = tree.GetValues(&a[0])
	if r == nil || *r[0] != b[0] || *r[1] != b[2] || len(r) != 2 {
		panic("Bad Put GetValues")
	}
}

func TestTreeMapPutReplace(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var a = [...]int{0, 0, 0, 1, 2, 3, 4, 5, 6, 7}
	var b = [...]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		tree.Put(&a[i], &b[i])
	}
	if tree.root == nil || tree.Size() != 8 {
		panic("Bad Put Replace")
	}
	var r = tree.GetValues(&a[0])
	if r == nil || *r[0] != b[0] || *r[1] != b[2] || len(r) != 2 {
		panic("Bad Put Replace")
	}
	tree.Replace(&a[0], &b[len(b)-1])
	r = tree.GetValues(&a[0])
	if r == nil || *r[0] != b[len(b)-1] || len(r) != 1 || tree.Size() != 8 {
		panic("Bad Put Replace")
	}
}

func TestTreeMapPutRemoveValue(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var i = tree.findInSlice(nil, nil)
	if i != -1 {
		panic("Bad Put RemoveValue")
	}
	var a = [...]int{0, 0, 0, 1, 1, 1, 4, 5, 6, 7}
	var b = [...]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		tree.Put(&a[i], &b[i])
	}
	if tree.root == nil || tree.Size() != 6 {
		panic("Bad Put RemoveValue")
	}
	var r = tree.GetValues(&a[0])
	if r == nil || *r[0] != b[0] || *r[1] != b[2] || len(r) != 2 {
		panic("Bad Put RemoveValue")
	}
	tree.RemoveValue(&a[0], &b[0])
	r = tree.GetValues(&a[0])
	if r == nil || *r[0] != b[2] || len(r) != 1 {
		panic("Bad Put RemoveValue")
	}

	tree.RemoveValue(&a[0], &b[2])
	r = tree.GetValues(&a[0])
	if r == nil || len(r) != 0 {
		panic("Bad Put RemoveValue")
	}

	tree.RemoveValue(&a[4], &b[4])
	r = tree.GetValues(&a[4])
	if r == nil || *r[0] != b[3] || *r[1] != b[5] || len(r) != 2 {
		panic("Bad Put RemoveValue")
	}
}

func TestTreeMapPutFind(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var i = 0
	tree.Put(&i, &i)
	var r = tree.Find(&i)
	if i != *r {
		panic("Bad Put Find")
	}
}

func TestTreeMapPutFindMin(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var r = tree.FindMin()
	if r != nil {
		panic("Bad Put FindMin")
	}
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		tree.Put(&a[i], &a[i])
	}
	r = tree.FindMin()
	if a[0] != *r {
		panic("Bad Put FindMin")
	}
}

func TestTreeMapPutRemoveMin(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		tree.Put(&a[i], &a[i])
	}
	if tree.root == nil || tree.Size() != 8 {
		panic("Bad Put RemoveValue")
	}
	tree.RemoveMin()
	if tree.root == nil || tree.Size() != 7 {
		panic("Bad Put RemoveValue")
	}
	var r = tree.Find(&a[0])
	if r != nil {
		panic("Bad Put RemoveMin")
	}
}

func TestTreeMapPutFindMax(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var r = tree.FindMax()
	if r != nil {
		panic("Bad Put FindMax")
	}
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		tree.Put(&a[i], &a[i])
	}
	r = tree.FindMax()
	if a[len(a)-1] != *r {
		panic("Bad Put FindMax")
	}
}

func TestTreeMapPutRemoveMax(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	for i := range a {
		tree.Put(&a[i], &a[i])
	}
	if tree.root == nil || tree.Size() != 8 {
		panic("Bad Put RemoveValue")
	}
	tree.RemoveMax()
	if tree.root == nil || tree.Size() != 7 {
		panic("Bad Put RemoveValue")
	}
	var r = tree.Find(&a[len(a)-1])
	if r != nil {
		panic("Bad Put RemoveMax")
	}
}

func TestTreeMapFindKeyValues(t *testing.T) {
	var tree = NewTreeMap[int, int]()
	var x, y = tree.FindKeyValues(nil)
	if x != nil || y != nil {
		panic("Bad Put FindKeyValues")
	}
	var a = [...]int{0, 0, 0, 1, 1, 1, 4, 5, 6, 7}
	var b = [...]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i := range a {
		tree.Put(&a[i], &b[i])
	}
	var k, r = tree.FindKeyValues(&a[0])
	if *k != a[0] || r == nil || *r[0] != b[0] || *r[1] != b[2] || len(r) != 2 {
		panic("Bad Put FindKeyValues")
	}
}

//!-
/* vim: tree tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
