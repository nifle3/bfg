package interpreter

import (
	"fmt"

	"github.com/nifle3/bfg/internal/core"
)

var _ core.Executor = &Executor{}

const stackLength int = 200_000_000

type Executor struct {
	stack      [stackLength]int
	curPointer int
}

func New() *Executor {
	return &Executor{
		stack:      [stackLength]int{},
		curPointer: 0,
	}
}

func (e *Executor) AddedOne() {
	e.stack[e.curPointer] += 1

}

func (e *Executor) LeftShift() {
	e.curPointer -= 1

	if e.curPointer < 0 {
		panic("index out of stack negative")
	}
}

func (e *Executor) PrintCurrent() {
	cur := e.stack[e.curPointer]
	char := rune(cur)
	fmt.Printf("%c", char)
}

func (e *Executor) ReduceOne() {
	e.stack[e.curPointer] -= 1
}

func (e *Executor) RightShift() {
	e.curPointer += 1

	if e.curPointer >= len(e.stack) {
		panic("index out of stack higher")
	}
}

func (e *Executor) ScanIntoCurrent() {
	var char rune
	fmt.Printf("\nPlease input one character: ")
	fmt.Scanf("%c", &char)
	e.stack[e.curPointer] = int(char)
}

func (e *Executor) StartWhile() {
}

func (e *Executor) StopWhile() {
}
