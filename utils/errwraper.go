package utils

import "fmt"

func WrapErr(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(format, a...)
	return fmt.Errorf("%s happen, errmsg:%s", msg, err.Error())
}
