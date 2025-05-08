package repl

import (
	"fmt"
	"iter"

	"github.com/nifle3/bfg/internal/core"
)

var _ core.Reader = Reader{}

type Reader struct {
}

func New() Reader {
	return Reader{}
}

func (r Reader) Read() iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		fmt.Println("Repl started!")
	inner:
		for {
			idx := 0
			var line string
			_, err := fmt.Scan(&line)
			if err != nil {
				panic("scan error")
			}

			for _, value := range line {
				if !yield(idx, value) {
					break inner
				}
				idx++
			}
		}
	}
}
