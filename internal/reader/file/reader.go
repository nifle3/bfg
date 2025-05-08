package file

import (
	"bufio"
	"iter"
	"os"

	"github.com/nifle3/bfg/internal/core"
)

var _ core.Reader = Reader{}

type Reader struct {
	filePath string
}

func New(filePath string) Reader {
	return Reader{
		filePath: filePath,
	}
}

func (r Reader) Read() iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		file, err := os.Open(r.filePath)
		if err != nil {
			panic("cannot open file")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

	inner:
		for scanner.Scan() {
			text := scanner.Text()
			for i, value := range text {
				if !yield(i, value) {
					break inner
				}
			}
		}
	}
}
