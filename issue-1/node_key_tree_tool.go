package avl

/*
 * This file was last modified at 2024.01.04 15:07 by Victor N. Skurikhin.
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org>
 * node_key_tree_tool.go
 * $Id$
 */
//!+

import (
	"cmp"
)

type nodeKey[K cmp.Ordered] struct {
	left, right *nodeKey[K]
	key         *K
	height      byte
}

// Вставка нового ключа в АВЛ-дерево выполняется, так же как это делается в простых деревьях поиска:
// спускаемся вниз по дереву, выбирая правое или левое направление движения в зависимости от результата сравнения ключа
// в текущем узле и вставляемого ключа. Единственное отличие заключается в том, что при возвращении из рекурсии
// (т.е. после того, как ключ вставлен либо в правое, либо в левое поддерево, и это дерево сбалансировано)
// выполняется балансировка текущего узла.
func insert[K cmp.Ordered](node *nodeKey[K], key *K) *nodeKey[K] {

	if key == nil {
		return nil
	}
	if node == nil {
		return &nodeKey[K]{key: key}
	}
	var c = cmp.Compare(*key, *node.key)
	if c == -1 { // key < node.key
		node.left = insert(node.left, key)
	} else if c == 1 { // key > node.key
		node.right = insert(node.right, key)
	} else { // key == node.key
		return node
	}
	return balance(node)
}

// Функция удаления ключа из АВЛ-дерева.
// Сначала находим нужный узел, выполняя те же действия, что и при вставке ключа.
func remove[K cmp.Ordered](node *nodeKey[K], key *K) (*nodeKey[K], bool) {

	if node == nil || key == nil {
		return nil, false
	}
	var found bool
	if cmp.Less(*key, *node.key) { // key < node.key
		node.left, found = remove(node.left, key)
	} else if cmp.Less(*node.key, *key) { // key > node.key
		node.right, found = remove(node.right, key)
	} else { // key == node.key

		var q = node.left
		var r = node.right
		// delete node; GC?
		if r == nil {
			return q, true
		}
		var minimal = findMin(r)
		minimal.right = removeMin(r)
		minimal.left = q

		return balance(minimal), true
	}
	return balance(node), found
}

// По свойству АВЛ-дерева у минимального элемента справа либо подвешен единственный узел, либо там пусто.
func removeMin[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	if node == nil {
		return nil
	}
	if node.left == nil {
		return node.right
	}
	node.left = removeMin(node.left)
	return balance(node)
}

// Функция находит минимальный ключ в этом поддереве.
// По свойству двоичного дерева поиска этот ключ находится в конце левой ветки, начиная от корня дерева.
func findMin[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

// По свойству АВЛ-дерева у максимального элемента слева либо подвешен единственный узел, либо там пусто.
func removeMax[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	if node == nil {
		return nil
	}
	if node.right == nil {
		return node.left
	}
	node.right = removeMax(node.right)

	return balance(node)
}

// Функция находит максимальный ключ в этом поддереве.
// По свойству двоичного дерева поиска этот ключ находится в конце правой ветки, начиная от корня дерева.
func findMax[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return findMax(node.right)
}

func find[K cmp.Ordered](node *nodeKey[K], key *K) *nodeKey[K] {

	if node == nil || key == nil {
		return nil
	}
	if cmp.Less(*key, *node.key) { // key < node.key
		return find(node.left, key)
	} else if cmp.Less(*node.key, *key) { // key > node.key
		return find(node.right, key)
	}
	return node // key == node.key
}

// Функция выполняющая балансировку, сводится к проверке условий и выполнению поворотов.
func balance[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	if node == nil {
		return nil
	}
	fixHeight(node)

	if balanceFactor(node) == 2 {
		if balanceFactor(node.right) < 0 {
			node.right = rotateRight(node.right)
		}
		return rotateLeft(node)
	}
	if balanceFactor(node) == -2 {
		if balanceFactor(node.left) > 0 {
			node.left = rotateLeft(node.left)
		}
		return rotateRight(node)
	}
	return node // балансировка не нужна
}

// Левый поворот является симметричной копией правого.
func rotateLeft[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	var up = node.right
	node.right = up.left
	up.left = node
	fixHeight(node)
	fixHeight(up)

	return up
}

// Функция реализующая правый поворот, выглядит следующим образом
// (как обычно, каждая функция, изменяющая дерево, возвращает новый корень полученного дерева).
func rotateRight[K cmp.Ordered](node *nodeKey[K]) *nodeKey[K] {

	var up = node.left
	node.left = up.right
	up.right = node
	fixHeight(node)
	fixHeight(up)

	return up
}

// Заметим, что все три функции являются нерекурсивными, т.е. время их работы есть величина О(1).
// Функция восстанавливает корректное значение поля height заданного узла
// (при условии, что значения этого поля в правом и левом дочерних узлах являются корректными).
func fixHeight[K cmp.Ordered](node *nodeKey[K]) {

	var hl = height(node.left)
	var hr = height(node.right)
	node.height = max(hl, hr) + 1
}

// Функция вычисляет balance factor заданного узла (и работает только с ненулевыми указателями).
func balanceFactor[K cmp.Ordered](node *nodeKey[K]) int {
	return int(height(node.right)) - int(height(node.left))
}

// Функция может работать и с нулевыми указателями (с пустыми деревьями).
func height[K cmp.Ordered](node *nodeKey[K]) byte {
	if node != nil {
		return node.height
	} else {
		return 0
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
