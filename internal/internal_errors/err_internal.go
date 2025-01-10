package internalerrors

import "net/http"

type ErrInternal struct {
	HttpError
	mensage string
}

func (e ErrInternal) Error() string {

	return e.mensage
}

func NewErrInternal(mensage string) ErrInternal {
	return ErrInternal{mensage: mensage, HttpError: HttpError{HttpStatus: http.StatusInternalServerError}}
}
