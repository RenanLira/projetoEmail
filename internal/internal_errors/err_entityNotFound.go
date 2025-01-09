package internalerrors

var ErrEntityNotFoundValue *ErrEntityNotFound

type ErrEntityNotFound struct {
	Entity string
}

func (e ErrEntityNotFound) Error() string {
	return e.Entity + " not found"
}

func NewErrEntityNotFound(entity string) ErrEntityNotFound {
	return ErrEntityNotFound{Entity: entity}
}
