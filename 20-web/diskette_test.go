package poker_test

import (
	"io"
	"poker"
	"testing"
)

func TestWriteDiskette(t *testing.T) {
	file, cleanUp := createTempFile(t, "12345")
	defer cleanUp()

	database := poker.NewDiskette(file)

	n, err := database.Write([]byte("abc"))
	if err != nil {
		t.Fatalf("unable to write to file: %v", err)
	}
	if n != 3 {
		t.Fatalf("expected to write %d bytes, but written %d", 3, n)
	}

	file.Seek(0, 0)
	buf, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("unable to read from file: %v", err)
	}

	got := string(buf)
	want := "abc"
	if got != want {
		t.Fatalf("got %q but want %q", got, want)
	}
}
