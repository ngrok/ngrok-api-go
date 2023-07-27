// Code generated for API Clients. DO NOT EDIT.

package ngrok

import (
	"fmt"
	"net/http"
)

// Returns true if the error is a not found response from the ngrok API.
func IsNotFound(err error) bool {
	if ee, ok := err.(*Error); ok {
		return int(ee.StatusCode) == http.StatusNotFound
	}
	return false
}

// Returns true if the given error is caused by any of the specified ngrok error codes.
// All ngrok error codes are documented at https://ngrok.com/docs/errors
func IsErrorCode(err error, codes ...int) bool {
	if ee, ok := err.(*Error); ok {
		for _, code := range codes {
			if ee.ErrorCode == fmt.Sprintf("ERR_NGROK_%d", code) {
				return true
			}
		}
	}
	return false
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Msg)
	if e.ErrorCode != "" {
		msg += fmt.Sprintf(" [%s]", e.ErrorCode)
	}
	if e.OperationID() != "" {
		msg += fmt.Sprintf("\n\nOperation ID: %s", e.OperationID())
	}
	return msg
}

// OperationID returns the unique trace ID assigned by ngrok to this API
// request.
func (e *Error) OperationID() string {
	return e.Details["operation_id"]
}
