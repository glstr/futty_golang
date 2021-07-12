package middleware

type CommonResponse struct {
	ErrorCode int32  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	RequestId int64  `json:"request_id"`
}
