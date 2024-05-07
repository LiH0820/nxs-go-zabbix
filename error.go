package zabbix

import "errors"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Is(target error) bool {
	var te Error
	if ok := errors.As(target, &te); !ok {
		return false
	}
	return e.Code == te.Code && e.Message == te.Message
}

func (e Error) IsOk() bool {
	return e.Code != 0
}

func newErr(code int, message, data string) error {
	return Error{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

var (
	ErrObjectNotExist = newErr(-32500, "Application error.", "No permissions to referred object or it does not exist!")
)
