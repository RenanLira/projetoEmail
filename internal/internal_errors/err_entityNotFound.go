package internalerrors

import "net/http"


type ErrEntityNotFound struct {
	HttpError
	Entity string
}

func (e ErrEntityNotFound) Error() string {
	return e.Entity + " not found"
}

func NewErrEntityNotFound(entity string) ErrEntityNotFound {
	return ErrEntityNotFound{Entity: entity, HttpError: HttpError{HttpStatus: http.StatusNotFound}}
}
