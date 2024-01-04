package avl

/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org>
 * tree_map.go
 * $Id$
 */
//!+

import (
	"cmp"
	"slices"
)

type TreeMap[K cmp.Ordered, V comparable] struct {
	root *nodeKey[K]
	hmap map[K][]*V
	size uint
}

func NewTreeMap[K cmp.Ordered, V comparable]() *TreeMap[K, V] {
	var treeMap = new(TreeMap[K, V])
	treeMap.hmap = make(map[K][]*V)
	return treeMap
}

func (t *TreeMap[K, V]) GetValues(key *K) []*V {

	if key != nil {
		return t.hmap[*key]
	}
	return nil
}

func (t *TreeMap[K, V]) Put(key *K, value *V) {

	if key != nil {
		if t.root = insert(t.root, key); t.root != nil {
			t.putValueToMap(key, value)
		}
	}
}

func (t *TreeMap[K, V]) putValueToMap(key *K, value *V) {

	if t.hmap[*key] == nil {
		t.hmap[*key] = []*V{value}
		t.size++
	} else {
		if i := t.findInSlice(t.hmap[*key], value); i > -1 {
			return
		}
		t.hmap[*key] = append(t.hmap[*key], value)
	}
}

func (t *TreeMap[K, V]) Replace(key *K, value *V) {

	if key != nil {
		if t.root = insert(t.root, key); t.root != nil {
			t.hmap[*key] = []*V{value}
		}
	}
}

func (t *TreeMap[K, V]) Remove(key *K) {

	if key != nil {
		var found bool
		t.root, found = remove(t.root, key)
		if found {
			delete(t.hmap, *key)
			t.size--
		}
	}
}

func (t *TreeMap[K, V]) RemoveValue(key *K, value *V) {

	if key != nil {
		var node = find(t.root, key)
		if node != nil {
			s := t.hmap[*key]
			i := t.findInSlice(s, value)
			t.hmap[*key] = t.removeFromSlice(s, i)
		}
	}
}

func (t *TreeMap[K, V]) findInSlice(slice []*V, value *V) int {
	if slice == nil {
		return -1
	}
	return slices.IndexFunc(slice, func(v *V) bool { return *v == *value })
}

func (t *TreeMap[K, V]) removeFromSlice(slice []*V, i int) []*V {
	return append(slice[:i], slice[i+1:]...)
}

func (t *TreeMap[K, V]) Find(key *K) *K {

	var node = find(t.root, key)
	if node == nil || key == nil {
		return nil
	} else {
		return node.key
	}
}

func (t *TreeMap[K, V]) FindKeyValues(key *K) (*K, []*V) {

	var node = find(t.root, key)
	if node == nil || key == nil {
		return nil, nil
	} else {
		return node.key, t.hmap[*node.key]
	}
}

func (t *TreeMap[K, V]) FindMin() *K {

	var minimal = findMin(t.root)
	if minimal == nil {
		return nil
	} else {
		return minimal.key
	}
}

func (t *TreeMap[K, V]) RemoveMin() {
	var r = t.root
	t.root = removeMin(t.root)
	if r != nil {
		t.size--
	}
}

func (t *TreeMap[K, V]) FindMax() *K {

	var maximal = findMax(t.root)
	if maximal == nil {
		return nil
	} else {
		return maximal.key
	}
}

func (t *TreeMap[K, V]) RemoveMax() {
	var r = t.root
	t.root = removeMax(t.root)
	if r != nil {
		t.size--
	}
}

func (t *TreeMap[K, V]) Size() uint {
	return t.size
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
