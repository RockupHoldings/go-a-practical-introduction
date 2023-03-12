package todo

import "errors"

var (
	ErrItemNotFound  = errors.New("item not found")
	ErrInvalidFilter = errors.New("invalid filter")
)

func IsNotFoundErr(err error) bool {
	return errors.Is(err, ErrItemNotFound)
}
