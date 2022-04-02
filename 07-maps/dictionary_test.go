package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("expected error, but none thrown")
	}
	if got.Error() != want.Error() {
		t.Errorf("expected %q, but got %q", want.Error(), got.Error())
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("no error expected, but thrown")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got, err := dictionary.Search("unknown")
		want := ""
		assertError(t, err, ErrNoItemDefined)
		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("replace existing", func(t *testing.T) {
		dictionary := Dictionary{"test": "original text"}
		err := dictionary.Add("test", "this is just a text")
		assertError(t, err, ErrItemExists)
		want := "original text"
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("update existing without error", func(t *testing.T) {
		dictionary := Dictionary{"test": "original text"}
		err := dictionary.Update("test", "this is just a text")
		assertNoError(t, err)
		want := "this is just a text"
		got, err := dictionary.Search("test")
		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("update non-existing with error", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is just a text")
		assertError(t, err, ErrNoItemToUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing without error", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a text"}
		err := dictionary.Delete("test")
		assertNoError(t, err)
		_, err = dictionary.Search("test")
		assertError(t, err, ErrNoItemDefined)
	})
	t.Run("delete non-existing with error", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")
		assertError(t, err, ErrNoItemToDelete)
	})
}
