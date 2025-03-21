package customErrors

type ApiError struct {
	Code int
	Err  error
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}
