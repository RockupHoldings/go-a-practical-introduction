package todo

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type FsStore struct {
	db *os.File
}

func NewFsStore(db *os.File) FsStore {
	return FsStore{db}
}

func (f FsStore) Create(task string) error {
	items, err := f.read()
	if err != nil {
		return err
	}

	uid := uuid.NewString()

	items = append(items, Item{
		Task:    task,
		UID:     uid,
		Checked: false,
	})

	return f.write(items)
}

func (f FsStore) Delete(uid string) error {
	items, err := f.read()
	if err != nil {
		return err
	}

	var index int
	var found bool

	for i, item := range items {
		if item.UID == uid {
			found = true
			index = i
			break
		}
	}

	if !found {
		return ErrItemNotFound
	}

	updatedItems := append(items[:index], items[index+1:]...)

	if err := f.write(updatedItems); err != nil {
		return err
	}

	return nil
}

func (f FsStore) List(filter Filter) ([]Item, error) {
	items, err := f.read()
	if err != nil {
		return nil, err
	}

	switch filter {
	case All:
		return items, nil
	case Checked:
		var itemsOut []Item
		for _, item := range items {
			if item.Checked {
				itemsOut = append(itemsOut, item)
			}
		}

		return itemsOut, nil
	case NotChecked:
		var itemsOut []Item
		for _, item := range items {
			if !item.Checked {
				itemsOut = append(itemsOut, item)
			}
		}
		return itemsOut, nil
	default:
		return nil, ErrInvalidFilter
	}
}

func (f FsStore) Toggle(uid string) error {
	items, err := f.read()
	if err != nil {
		return err
	}

	var found bool
	for i, item := range items {
		if item.UID == uid {
			found = true
			items[i].Checked = !item.Checked
			break
		}
	}

	if !found {
		return ErrItemNotFound
	}

	if err := f.write(items); err != nil {
		return err
	}

	return nil
}

func (f FsStore) read() ([]Item, error) {
	f.db.Seek(0, 0)

	var out []Item
	err := json.NewDecoder(f.db).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (f FsStore) write(items []Item) error {
	f.db.Truncate(0)
	f.db.Seek(0, 0)

	err := json.NewEncoder(f.db).Encode(items)
	if err != nil {
		return err
	}

	return nil
}
