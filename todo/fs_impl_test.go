package todo_test

import (
	"encoding/json"
	"os"
	"rock_ed/todo"
	"testing"
)

func TestTodoFS(t *testing.T) {
	initialData, err := json.Marshal([]todo.Item{
		{
			Task:    "Go to grocery store",
			UID:     "1",
			Checked: false,
		},
		{
			Task:    "Do homework",
			UID:     "2",
			Checked: false,
		},
	})
	isNilErr(t, err)

	db, cleanup := createTempFile(t, string(initialData))
	defer cleanup()

	todoStore := todo.NewFsStore(db)

	t.Run("create and list items", func(t *testing.T) {
		isNilErr(t, todoStore.Create("one"))
		isNilErr(t, todoStore.Create("two"))

		items, err := todoStore.List(todo.All)
		isNilErr(t, err)

		if len(items) != 4 {
			t.Errorf("expected len(items)=4; but got %d", len(items))
		}
	})

	t.Run("delete item", func(t *testing.T) {
		err := todoStore.Delete("does_not_exist")

		if !todo.IsNotFoundErr(err) {
			t.Errorf("expected ErrItemNotFound, but got %q", err)
		}

		err = todoStore.Delete("1")
		isNilErr(t, err)

		items, err := todoStore.List(todo.All)
		isNilErr(t, err)
		if len(items) != 3 {
			t.Errorf("expected len(items)=3; but got %d", len(items))
		}
	})

	t.Run("toggle item", func(t *testing.T) {
		if err := todoStore.Toggle("does_not_exist"); !todo.IsNotFoundErr(err) {
			t.Errorf("expected ErrItemNotFound, but got %q", err)
		}

		err := todoStore.Toggle("2")
		isNilErr(t, err)

		items, err := todoStore.List(todo.All)
		isNilErr(t, err)

		var found bool
		for _, item := range items {
			if item.UID == "2" {
				found = true
				if !item.Checked {
					t.Errorf("expected checked items, but got unchecked")
				}
			}
		}

		if !found {
			t.Fatal("item not found")
		}
	})

	t.Run("filter", func(t *testing.T) {
		items, err := todoStore.List(todo.Checked)
		isNilErr(t, err)

		if len(items) != 1 {
			t.Errorf("expected len(items)=1; but got %d", len(items))
		}

		items, err = todoStore.List(todo.NotChecked)
		isNilErr(t, err)

		if len(items) != 2 {
			t.Errorf("expected len(items)=2; but got %d", len(items))
		}
	})
}

func isNilErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected nil err, but got '%s'", err)
	}
}

// file_system_store_test.go
func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	if _, err := tmpfile.Write([]byte(initialData)); err != nil {
		t.Fatal(err)
	}

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
