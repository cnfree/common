package common

import (
	"sync"
	"github.com/cnfree/common/pool"
	"time"
)

type poolUtil struct {
	mutex sync.Mutex
}

var Pool = poolUtil{}

func (this *poolUtil) NewGoRoutinePool(maxGoRoutineNum int) *pool.GoRoutinePool {
	return pool.NewGoRoutinePool(maxGoRoutineNum)
}

func (this *poolUtil) NewBufferPool(maxBufferNum, initBufferSize int) *pool.BufferPool {
	return pool.NewBufferPool(maxBufferNum, initBufferSize)
}

func (this *poolUtil) NewObjectPool(maxObjectNum int, createObj func() interface{}, resetObj func(interface{})) *pool.ObjectPool {
	return pool.NewObjectPool(maxObjectNum, createObj, resetObj)
}

func (this *poolUtil) NewResourcePool(factory pool.ResourceFactory, capacity, maxCapacity int, idleTimeout time.Duration) *pool.ResourcePool {
	return pool.NewResourcePool(factory, capacity, maxCapacity, idleTimeout)
}
