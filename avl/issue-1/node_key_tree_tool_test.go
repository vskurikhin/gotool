package avl

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	exitCode := m.Run()
	// Exit
	os.Exit(exitCode)
}

func TestHeight(t *testing.T) {
	var node = &nodeKey[int]{height: 13}
	var test = height(node)
	if test != 13 {
		panic("Bad height")
	}
	test = height[int](nil)
	if test != 0 {
		panic("Bad height")
	}
}

func TestBalanceFactor(t *testing.T) {
	var nodeLeft = &nodeKey[int]{height: 2}
	var nodeRight = &nodeKey[int]{height: 1}
	var node = &nodeKey[int]{left: nodeLeft, right: nodeRight}
	var test = balanceFactor(node)
	if test != -1 {
		panic("Bad height")
	}
	node = &nodeKey[int]{}
	test = balanceFactor[int](node)
	if test != 0 {
		panic("Bad height")
	}
}

func TestFixHeight(t *testing.T) {
	var nodeLeft = &nodeKey[int]{height: 2}
	var nodeRight = &nodeKey[int]{height: 1}
	var node = &nodeKey[int]{left: nodeLeft, right: nodeRight}
	fixHeight(node)
	if node.height != 3 {
		panic("Bad height")
	}
}

func TestRotateRight(t *testing.T) {
	var i1 = 1
	var i2 = 2
	var i3 = 3
	var nodeLeft = &nodeKey[int]{key: &i1, height: 1}
	var nodeRight = &nodeKey[int]{key: &i3, height: 1}
	var node = &nodeKey[int]{key: &i2, left: nodeLeft, right: nodeRight}
	var up = rotateRight(node)
	if up != nodeLeft {
		panic("Bad rotateRight")
	}
	if up.right != node {
		panic("Bad rotateRight")
	}
	if up.right.right != nodeRight {
		panic("Bad rotateRight")
	}
}

func TestRotateLeft(t *testing.T) {
	var i1 = 1
	var i2 = 2
	var i3 = 3
	var nodeLeft = &nodeKey[int]{key: &i1, height: 1}
	var nodeRight = &nodeKey[int]{key: &i3, height: 1}
	var node = &nodeKey[int]{key: &i2, left: nodeLeft, right: nodeRight}
	var up = rotateLeft(node)
	if up != nodeRight {
		panic("Bad rotateLeft")
	}
	if up.left != node {
		panic("Bad rotateLeft")
	}
	if up.left.left != nodeLeft {
		panic("Bad rotateLeft")
	}
}

func TestBalanceCase0(t *testing.T) {
	var i0 = 0
	balance[int](nil)
	var a = &nodeKey[int]{key: &i0, height: 1}
	var up = balance(a)
	if up == nil {
		panic("Bad balance")
	}
	if up != a {
		panic("Bad balance")
	}
	if up.height != 1 {
		panic("Bad balance")
	}
}

func TestBalanceCase1(t *testing.T) {
	var i = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//var tree = NewTreeSet[int]()
	//for j := range i {
	//	tree.Put(&j)
	//}
	//
	var a = &nodeKey[int]{key: &i[1], height: 1}
	fixHeight(a)
	var b = &nodeKey[int]{key: &i[3], height: 1}
	fixHeight(b)
	var c = &nodeKey[int]{key: &i[5], height: 1}
	fixHeight(c)
	var d = &nodeKey[int]{key: &i[7], height: 1}
	fixHeight(d)
	var s = &nodeKey[int]{key: &i[4], left: b, right: c}
	fixHeight(s)
	var q = &nodeKey[int]{key: &i[6], left: s, right: d}
	fixHeight(q)
	var p = &nodeKey[int]{key: &i[2], left: a, right: q}
	fixHeight(p)
	var up = balance(p)
	if up == nil {
		panic("Bad balance")
	}
	if up != s {
		panic("Bad balance")
	}
	if up.height != 3 {
		panic("Bad balance")
	}
	if up.left != p {
		panic("Bad balance")
	}
	if p.height != 2 {
		panic("Bad balance")
	}
	if p.left != a {
		panic("Bad balance")
	}
	if a.height != 1 {
		panic("Bad balance")
	}
	if p.right != b {
		panic("Bad balance")
	}
	if b.height != 1 {
		panic("Bad balance")
	}
	if up.right != q {
		panic("Bad balance")
	}
	if q.height != 2 {
		panic("Bad balance")
	}
	if q.left != c {
		panic("Bad balance")
	}
	if c.height != 1 {
		panic("Bad balance")
	}
	if q.right != d {
		panic("Bad balance")
	}
	if d.height != 1 {
		panic("Bad balance")
	}
}

func TestBalanceCase2(t *testing.T) {
	var i = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var a = &nodeKey[int]{key: &i[1], height: 1}
	fixHeight(a)
	var b = &nodeKey[int]{key: &i[3], height: 1}
	fixHeight(b)
	var c = &nodeKey[int]{key: &i[5], height: 1}
	fixHeight(c)
	var d = &nodeKey[int]{key: &i[7], height: 1}
	fixHeight(d)
	var s = &nodeKey[int]{key: &i[4], left: b, right: c}
	fixHeight(s)
	var q = &nodeKey[int]{key: &i[2], left: a, right: s}
	fixHeight(q)
	var p = &nodeKey[int]{key: &i[6], left: q, right: d}
	fixHeight(p)
	var up = balance(p)
	if up == nil {
		panic("Bad balance")
	}
	if up != s {
		panic("Bad balance")
	}
	if up.height != 3 {
		panic("Bad balance")
	}
	if up.left != q {
		panic("Bad balance")
	}
	if p.height != 2 {
		panic("Bad balance")
	}
	if q.left != a {
		panic("Bad balance")
	}
	if a.height != 1 {
		panic("Bad balance")
	}
	if q.right != b {
		panic("Bad balance")
	}
	if b.height != 1 {
		panic("Bad balance")
	}
	if up.right != p {
		panic("Bad balance")
	}
	if p.height != 2 {
		panic("Bad balance")
	}
	if p.left != c {
		panic("Bad balance")
	}
	if c.height != 1 {
		panic("Bad balance")
	}
	if p.right != d {
		panic("Bad balance")
	}
	if d.height != 1 {
		panic("Bad balance")
	}
}

func TestInsertCase0(t *testing.T) {
	var a = [...]int{0, 0}
	var root *nodeKey[int]
	root = insert[int](root, nil)
	for i := range a {
		root = insert[int](root, &a[i])
	}
	if root == nil && root.height != 1 && *root.key != 0 {
		panic("Bad insert")
	}
	if root.left != nil && root.right != nil {
		panic("Bad insert")
	}
}

func TestInsertCase1(t *testing.T) {
	var a = [...]int{9, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	for i := range a {
		root = insert[int](root, &a[i])
	}
	if root == nil && root.height != 3 && *root.key != 4 {
		panic("Bad insert")
	}
	if root.left == nil && root.left.height != 2 && *root.left.key != 2 {
		panic("Bad insert")
	}
	if root.left.left == nil && root.left.left.height != 1 && *root.left.left.key != 1 {
		panic("Bad insert")
	}
	if root.left.right == nil && root.left.right.height != 1 && *root.left.right.key != 3 {
		panic("Bad insert")
	}
	if root.right == nil && root.right.height != 2 && *root.right.key != 6 {
		panic("Bad insert")
	}
	if root.right.left == nil && root.right.left.height != 1 && *root.right.left.key != 5 {
		panic("Bad insert")
	}
	if root.right.right == nil && root.right.right.height != 1 && *root.right.right.key != 9 {
		panic("Bad insert")
	}
}

func TestRemoveCase0(t *testing.T) {
	var a = [...]int{0, 0}
	var root *nodeKey[int]
	var found bool
	root, found = remove[int](nil, nil)
	if root != nil || found {
		panic("Bad remove")
	}
	for i := range a {
		root = insert[int](root, &a[i])
	}
	var i0 = 0
	root, found = remove(root, &i0)
	if root != nil || !found {
		panic("Bad remove")
	}
}

func TestRemoveCase1(t *testing.T) {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	for i := range a {
		root = insert[int](root, &a[i])
	}
	for i := range a {
		root, _ = remove[int](root, &a[i])
	}
	if root != nil {
		panic("Bad remove")
	}
	for i := range a {
		root = insert[int](root, &a[i])
	}
	for i := len(a) - 1; i > -1; i-- {
		root, _ = remove[int](root, &a[i])
	}
	if root != nil {
		panic("Bad remove")
	}
}

func TestFindMinCase1(t *testing.T) {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	root = findMin(root)
	for i := range a {
		root = insert[int](root, &a[i])
	}
	minimal := findMin(root)
	if minimal == nil || *minimal.key != 0 {
		panic("Bad findMin")
	}
}

func TestRemoveMinCase1(t *testing.T) {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	root = removeMin(root)
	for i := range a {
		root = insert[int](root, &a[i])
	}
	root = removeMin(root)
	if root == nil || *root.key != 3 {
		panic("Bad removeMin")
	}
}

func TestFindMaxCase1(t *testing.T) {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	root = findMax(root)
	for i := range a {
		root = insert[int](root, &a[i])
	}
	maximal := findMax(root)
	if maximal == nil || *maximal.key != 7 {
		panic("Bad findMin")
	}
}

func TestRemoveMaxCase1(t *testing.T) {
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	root = removeMax(root)
	for i := range a {
		root = insert[int](root, &a[i])
	}
	root = removeMax(root)
	if root == nil || *root.key != 3 {
		panic("Bad removeMin")
	}
}

func TestFindCase0(t *testing.T) {
	var a = [...]int{0, 0}
	var root *nodeKey[int]
	root = find[int](nil, nil)
	if root != nil {
		panic("Bad remove")
	}
	for i := range a {
		root = insert[int](root, &a[i])
	}
	root = find(root, nil)
	if root != nil {
		panic("Bad remove")
	}
}

func TestFindCase1(t *testing.T) {
	var a = [...]int{9, 1, 2, 3, 4, 5, 6, 7}
	var root *nodeKey[int]
	for i := range a {
		root = insert[int](root, &a[i])
	}
	var i7 = 7
	result := find(root, &i7)
	if result == nil || *result.key != 7 {
		panic("Bad find")
	}
}

//!-
/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
