package core

import "iter"

type Reader interface {
	Read() iter.Seq2[int, rune]
}
