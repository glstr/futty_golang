package errcode

type ErrorInfo struct {
	ErrorCode int32  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

var (
	Ok            ErrorInfo = ErrorInfo{0, "ok"}
	InternalError ErrorInfo = ErrorInfo{1, "internal error"}
)
