package serror

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrIs系统异常 = &Error{
		Code:    101,
		Message: "系统异常",
	}

	ErrIs请求状态异常 = &Error{
		Code:    102,
		Message: "http状态异常",
	}

	ErrIs数据解析异常 = &Error{
		Code:    103,
		Message: "数据解析异常",
	}
)

// New 业务异常
func New(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
