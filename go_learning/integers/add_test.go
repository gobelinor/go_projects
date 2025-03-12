package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("1+1", func(t *testing.T) {
		got := Add(1, 1)
		want := 2
		assertCorrectInt(t, got, want)
	})
}

func assertCorrectInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
