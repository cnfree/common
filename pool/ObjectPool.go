package pool

import "sync"

// CreateFunc is used by ObjectPool to create a new object when it's empty.
type CreateFunc func() interface{}

// ResetFunc is used by ObjectPool to reset a used object to it's initial state for reuse.
type ResetFunc func(interface{})

// NewObjectPool is the only way to get a new, ready-to-use ObjectPool for objects of a specified type.
//
// If you use `var op pool.ObjectPool`, or `new(pool.ObjectPool)`, or the like to obtain an ObjectPool, it'll
// crash when you call Get().
//
//   maxObjectNum: Maximum number of objects that will be pooled in ObjectPool.
//   createObj: Called to create a new object when ObjectPool is empty. Cannot be nil.
//   clearObj: Called to reset a used object to it's initial state for reuse. Could be nil if it need not to be reset.
//             ObjectPool will perform about 12% faster if `clearObj` is nil.
//
// Example:
//
//   // create an ObjectPool for bytes.Buffer
//   op := pool.NewObjectPool(10000,
//                            func() interface{} { return new(bytes.Buffer) },
//                            func(obj interface{}) { obj.(*bytes.Buffer).Reset() })
//   obj := op.Get() // get a ready-to-use bytes.Buffer
//   buf := obj.(*bytes.Buffer)
//   // do something with `buf`
//   op.Put(obj) // return obj to ObjectPool. `op.Put(buf)` is OK too.
func NewObjectPool(maxObjectNum int, createObj CreateFunc, resetObj ResetFunc) *ObjectPool {
	return &ObjectPool{maxObjNum: maxObjectNum, createFunc: createObj, resetFunc: resetObj}
}

// ObjectPool is a goroutine-safe generic pool for objects of any type.
// It performs about 36% slower than the other specialized object pools such as BufferPool.
type ObjectPool struct {
	lock       sync.Mutex
	freeList   *object
	freeObjNum int
	maxObjNum  int
	createFunc CreateFunc
	resetFunc  ResetFunc
}

// object holds an object of arbitrary type for reuse.
type object struct {
	obj  interface{}
	next *object
}

// Get returns a ready-to-use object.
func (op *ObjectPool) Get() interface{} {
	op.lock.Lock()
	o := op.freeList
	if o != nil {
		op.freeList = o.next
		op.freeObjNum--
	}
	op.lock.Unlock()

	var obj interface{}
	if o != nil {
		obj = o.obj
		if op.resetFunc != nil {
			op.resetFunc(obj)
		}
		o.obj = nil
		o.next = nil
	} else {
		obj = op.createFunc()
	}
	return obj
}

// Put returns an object to ObjectPool.
func (op *ObjectPool) Put(obj interface{}) {
	op.lock.Lock()
	if op.freeObjNum < op.maxObjNum {
		op.freeList = &object{obj, op.freeList}
		op.freeObjNum++
	}
	op.lock.Unlock()
}
