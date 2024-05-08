package model

type ErrNotFound struct {
	error string
}

func (e *ErrNotFound) Error() string {
	return e.error
}
