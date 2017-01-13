package common

import (
	"sync"
	"github.com/cnfree/common/collection"
	"github.com/cnfree/common/cmp"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/maps/hashbidimap"
	"github.com/emirpasic/gods/maps/treebidimap"
	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/emirpasic/gods/trees/binaryheap"
)

type collectionUtil struct {
	mutex sync.Mutex
}

var Collection = collectionUtil{}

func (this collectionUtil) NewLinkedOrderedMap() *collection.LinkedOrderedMap {
	return collection.NewLinkedOrderedMap(cmp.Compare)
}

func (this collectionUtil) NewLockfreeQueue() *collection.LockfreeQueue {
	return collection.NewLockfreeQueue()
}

func (this collectionUtil) NewArrayList() *arraylist.List {
	return arraylist.New()
}

func (this collectionUtil) NewSinglyLinkedList() *singlylinkedlist.List {
	return singlylinkedlist.New()
}

func (this collectionUtil) NewDoublyLinkedList() *doublylinkedlist.List {
	return doublylinkedlist.New()
}

func (this collectionUtil) NewHashSet() *hashset.Set {
	return hashset.New()
}

func (this collectionUtil) NewTreeSet() *treeset.Set {
	return treeset.NewWith(cmp.Compare)
}

func (this collectionUtil) NewTreeSetWithComparator(comparator cmp.Comparator) *treeset.Set {
	return treeset.NewWith(utils.Comparator(comparator))
}

func (this collectionUtil) NewLinkedListStack() *linkedliststack.Stack {
	return linkedliststack.New()
}

func (this collectionUtil) NewArrayStack() *arraystack.Stack {
	return arraystack.New()
}

func (this collectionUtil) NewHashMap() *hashmap.Map {
	return hashmap.New()
}

func (this collectionUtil) NewTreeMap() *treemap.Map {
	return treemap.NewWith(cmp.Compare)
}

func (this collectionUtil) NewTreeMapWithComparator(comparator cmp.Comparator) *treemap.Map {
	return treemap.NewWith(utils.Comparator(comparator))
}

func (this collectionUtil) NewHashBidiMap() *hashbidimap.Map {
	return hashbidimap.New()
}

func (this collectionUtil) NewTreeBidiMap() *treebidimap.Map {
	return treebidimap.NewWith(cmp.Compare, cmp.Compare)
}

func (this collectionUtil) NewTreeBidiMapWithComparator(keyComparator cmp.Comparator, valueComparator cmp.Comparator) *treebidimap.Map {
	return treebidimap.NewWith(utils.Comparator(keyComparator), utils.Comparator(valueComparator))
}

func (this collectionUtil) NewRedBlackTree() *redblacktree.Tree {
	return redblacktree.NewWith(cmp.Compare)
}

func (this collectionUtil) NewRedBlackTreeWithComparator(comparator cmp.Comparator) *redblacktree.Tree {
	return redblacktree.NewWith(utils.Comparator(comparator))
}

func (this collectionUtil) NewBTree(order int) *btree.Tree {
	return btree.NewWith(order, cmp.Compare)
}

func (this collectionUtil) NewBTreeWithComparator(order int, comparator cmp.Comparator) *btree.Tree {
	return btree.NewWith(order, utils.Comparator(comparator))
}

func (this collectionUtil) NewBinaryHeap() *binaryheap.Heap {
	return binaryheap.NewWith(cmp.Compare)
}

func (this collectionUtil) NewBinaryHeapWithComparator(order int, comparator cmp.Comparator) *binaryheap.Heap {
	return binaryheap.NewWith(utils.Comparator(comparator))
}
