package maths

const (
	MaxInt      = 1<<(IntBits-1) - 1
	MinInt      = -MaxInt - 1
	MaxUint     = 1<<IntBits - 1
	IntBits     = 1 << (^uint(0)>>32&1 + ^uint(0)>>16&1 + ^uint(0)>>8&1 + 3)
	UintPtrBits = 1 << (^uintptr(0)>>32&1 + ^uintptr(0)>>16&1 + ^uintptr(0)>>8&1 + 3)
)
