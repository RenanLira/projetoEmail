package internalerrors

import "net/http"

type ErrUnauthorized struct {
	HttpError
}

func (e ErrUnauthorized) Error() string {
	return "unauthorized"
}

func NewErrUnauthorized() ErrUnauthorized {
	return ErrUnauthorized{
		HttpError: HttpError{HttpStatus: http.StatusUnauthorized},
	}
}
