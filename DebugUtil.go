package common

import (
	"fmt"
	"sync"
	"io"
	"github.com/cnfree/common/debug"
)

type debugUtil struct {
	mutex sync.Mutex
}

var Debug = debugUtil{}

func (this debugUtil) StackInfo(all bool) string {
	return debug.StackInfo(all)
}

func (this debugUtil) PrintStack(all bool) {
	fmt.Println("Current stack is: ", this.StackInfo(all))
}

func (this debugUtil) ColorPrint(w io.Writer, s string, c debug.ConsoleColor) {
	debug.ColorPrint(w, s, c)
}

func (this debugUtil) Trace() debug.Stack {
	return debug.Trace()
}

func (this debugUtil) TraceN(skip, depth int) debug.Stack {
	return debug.TraceN(skip, depth)
}
