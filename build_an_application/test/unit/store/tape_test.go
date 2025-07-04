package store_test

import (
	"io"
	"testing"

	"example.com/build_an_application/src/store"
)

func TestTape_Write(t *testing.T) {
	file, clean := store.CreateTempFile(t, "12345")
	defer clean()

	tape := &store.Tape{File: file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
