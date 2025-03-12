package main

import "testing"

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func asserteq(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestSearch(t *testing.T) {
	t.Run("search basic", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		asserteq(t, got, want)
	})
	t.Run("error basic", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("gougougaga")
		want := ErrNotFound
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, want)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	asserteq(t, got, definition)
}

func TestAdd(t *testing.T) {
	t.Run("basic add", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("add existing", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("basic Update", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		dictionary.Add(word, definition)
		definition2 := "ouga bouga test"
		err := dictionary.Update(word, definition2)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition2)
	})
	t.Run("update non existing", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition2 := "ouga bouga test"
		err := dictionary.Update(word, definition2)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("basic delete", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		dictionary.Add(word, definition)
		err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("delete non existing", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
