package internalerrors

type HttpErrorImp interface {
	GetStatus() int
	error
}


type HttpError struct {
	HttpStatus int
}


func (e HttpError) GetStatus() int {
	return e.HttpStatus
}
