// Package albresponse - Define functions for handling common 5xx Status Codes
package albresponse

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Default message for an InternalServerError Status Code response
var internalServerErrorMessage = "Internal Server Error"

// Default message for a NotImplemented Status Code response
var notImplementedMessage = "Not implemented"

// InternalServerError - Return a 500 Status code with a message
func InternalServerError(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = internalServerErrorMessage
	}
	return errorResponse(http.StatusInternalServerError, message)
}

// NotImplemented - Return a 501 Status code with a message
func NotImplemented(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = notImplementedMessage
	}
	return errorResponse(http.StatusNotImplemented, message)
}
