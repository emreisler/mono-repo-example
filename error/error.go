package error

type ErrorCode int
type HttpStatus int

type MonoError struct {
	Timestamp    int64
	ErrorCode    ErrorCode
	HttpStatus   HttpStatus
	ErrorMessage string
	Err          error
}
