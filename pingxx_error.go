package pingxx_error

import (
	"fmt"
	"runtime"
	"strings"
)

type PingxxErr struct {
	Code    int
	Message string
	DeepErr error //底层错误信息 可能为nil

	Filename string
	Line     int
}

func New(code int, message string, err interface{}) *PingxxErr {
	pingxxErr := &PingxxErr{}
	switch t := err.(type) {
	case PingxxErr:
		return &t
	case *PingxxErr:
		return t
	case error:
		pingxxErr.DeepErr = t
	default:
		pingxxErr.DeepErr = fmt.Errorf("%v", t)
	}
	pingxxErr.Code = code
	pingxxErr.Message = message

	_, file, line, ok := runtime.Caller(1)
	if ok {
		pingxxErr.Line = line
		components := strings.Split(file, "/")
		pingxxErr.Filename = components[(len(components) - 1)]
	}
	return pingxxErr
}

func (this *PingxxErr) Error() string {
	str := fmt.Sprintf("发生错误地点[%s:%d] Code[%d] Message[%s] ", this.Filename, this.Line, this.Code, this.Message)
	if this.DeepErr != nil {
		str = fmt.Sprintf("%s 底层错误信息 %s", str, this.DeepErr.Error())
	}
	return str
}
