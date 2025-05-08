package core

type Executor interface {
	LeftShift()
	RightShift()
	AddedOne()
	ReduceOne()
	PrintCurrent()
	ScanIntoCurrent()
	StartWhile()
	StopWhile()
}
