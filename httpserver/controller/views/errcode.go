package views

import (
	"errors"
	"net/http"

	"github.com/glstr/futty_golang/model"
)

var (
	ErrParamInvalid = errors.New("param error")
	ErrInternalErr  = errors.New("internal error")
)

var ErrStatusMap = map[error]int{
	ErrParamInvalid: http.StatusBadRequest,

	model.ErrNotFound: http.StatusNotFound,
}

func GetErrInfoFromErr(err error) (int32, string) {
	if err == nil {
		return 0, "success"
	}

	if errCode, ok := ErrStatusMap[err]; ok {
		return int32(errCode), err.Error()
	}

	return http.StatusInternalServerError, ErrInternalErr.Error()
}
