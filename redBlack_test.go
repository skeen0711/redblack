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

	t.Run("multiple insertions", func(t *testing.T) {
		tree := NewRedBlackTree()
		tree.Insert(1, "one")
		tree.Insert(2, "two")
		tree.Insert(3, "three")

		got, found := tree.Search(1)
		if !found || got != "one" {
			t.Errorf("expected to find key 1, found %q", got)
		}
		got, found = tree.Search(3)
		if !found || got != "three" {
			t.Errorf("expected to find key 3, found %q", got)
		}
	})

	t.Run("Search returns empty string and false if key not found", func(t *testing.T) {
		tree := NewRedBlackTree()
		tree.Insert(1, "one")
		tree.Insert(2, "two")
		tree.Insert(3, "three")
		tree.Insert(5, "five")

		_, found := tree.Search(4)
		if found {
			t.Errorf("Key: 4 was found but was not in tree")
		}
	})
}
