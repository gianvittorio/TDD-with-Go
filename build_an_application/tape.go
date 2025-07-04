package poker

import (
	"io"
	"os"
)

type Tape struct {
	File *os.File
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.File.Truncate(0)
	t.File.Seek(0, io.SeekStart)

	return t.File.Write(p)
}
