package internalerrors

type ErrInternal struct {
	mensage string
}

func (e ErrInternal) Error() string {
	return e.mensage
}

func NewErrInternal(mensage string) ErrInternal {
	return ErrInternal{mensage: mensage}
}