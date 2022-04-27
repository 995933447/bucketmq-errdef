package errdef

import "fmt"

type Err struct {
	code int32
	msg string
}
func NewErr(code int32, msg string) *Err {
	return &Err{
		code: code,
		msg: msg,
	}
}

func (e *Err) Error() string {
	return fmt.Sprintf("error code:%d, error msg:%s", e.code, e.msg)
}

func (e *Err) GetCode() int32 {
	return e.code
}

func (e *Err) GetMsg() string {
	return e.msg
}
