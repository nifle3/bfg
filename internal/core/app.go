package core

func Run(
	executor Executor,
	reader Reader,
) {
	operatorMap := map[rune]func(){
		'>': executor.RightShift,
		'<': executor.LeftShift,
		'+': executor.AddedOne,
		'-': executor.ReduceOne,
		',': executor.ScanIntoCurrent,
		'.': executor.PrintCurrent,
		'[': executor.StartWhile,
		']': executor.StopWhile,
	}

	for _, value := range reader.Read() {
		operator, ok := operatorMap[value]
		if !ok {
			continue
		}

		operator()
	}
}
