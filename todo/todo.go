package todo

// notice that we avoid stutter when deciding on a name, the usage
// for this interface is todo.Store and not todo.TodoStore
type Store interface {
	Create(string) error
	Delete(string) error
	List(Filter) ([]Item, error)
	Toggle(string) error
}

type Filter int

const (
	// Filters are generally truthy, meaning when we pass Checked, it means we
	// keep all checked Items
	All Filter = iota
	Checked
	NotChecked
)

type Item struct {
	Task    string `json:"task"`
	UID     string `json:"uid"`
	Checked bool   `json:"checked"`
}
