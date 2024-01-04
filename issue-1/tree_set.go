package avl

/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org>
 * tree_set.go
 * $Id$
 */
//!+

import (
	"cmp"
)

type TreeSet[K cmp.Ordered] struct {
	root *nodeKey[K]
	size uint
}

func NewTreeSet[K cmp.Ordered]() *TreeSet[K] {
	return new(TreeSet[K])
}

func (t *TreeSet[K]) Put(key *K) {
	if key != nil {
		t.root = insert(t.root, key)
		t.size++
	}
}

func (t *TreeSet[K]) Remove(key *K) {
	if key != nil {
		var found bool
		t.root, found = remove(t.root, key)
		if found {
			t.size--
		}
	}
}

func (t *TreeSet[K]) Find(key *K) *K {

	var node = find(t.root, key)
	if node == nil || key == nil {
		return nil
	} else {
		return node.key
	}
}

func (t *TreeSet[K]) FindMin() *K {

	var minimal = findMin(t.root)
	if minimal == nil {
		return nil
	} else {
		return minimal.key
	}
}

func (t *TreeSet[K]) RemoveMin() {
	var r = t.root
	t.root = removeMin(t.root)
	if r != nil {
		t.size--
	}
}

func (t *TreeSet[K]) FindMax() *K {

	var maximal = findMax(t.root)
	if maximal == nil {
		return nil
	} else {
		return maximal.key
	}
}

func (t *TreeSet[K]) RemoveMax() {
	var r = t.root
	t.root = removeMax(t.root)
	if r != nil {
		t.size--
	}
}

func (t *TreeSet[K]) Size() uint {
	return t.size
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
