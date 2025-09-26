package redblack

import "testing"

func TestRedBlackTree(t *testing.T) {
	t.Run("insert and search single value", func(t *testing.T) {
		tree := NewRedBlackTree()
		tree.Insert(1, "one")

		got, found := tree.Search(1)
		if !found {
			t.Errorf("expected to find key 1")
		}
		if got != "one" {
			t.Errorf("got %q want %q", got, "one")
		}
	})
}
