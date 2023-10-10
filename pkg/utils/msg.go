package utils

import "strings"
import "errors"

func SubstrError(err error) (sErr error) {
	start := strings.Index(err.Error(), ":")
	msg := err.Error()[start+2:]
	sErr = errors.New(msg)
	return
}
