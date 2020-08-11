// Package albresponse - Define functions for handling common 4xx Status Codes
package albresponse

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Default message for a BadRequest Status Code response
var badRequestMessage = "InvalidRequest"

// Default message for an Unauthorized Status Code response
var unauthorizedMessage = "Unauthorized"

// Default message for a Forbidden Status Code response
var forbiddenMessage = "Forbidden"

// Default message for a NotFound Status Code response
var notFoundMessage = "Not Found"

// Default message for a MethodNotAllowed Status Code response
var methodNotAllowedMessage = "Method not allowed"

// Default message for a NotAcceptable Status Code response
var notAcceptableMessage = "Not acceptable"

// Default message for a PreconditionFailed Status Code response
var preconditionFailedMessage = "Precondition Failed"

// Default message for a UnsupportedMediaType Status Code response
var unsupportedMediaTypeMessage = "Unsupported Media Type"

// BadRequest - Return a 400 Status Code with a message
func BadRequest(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = badRequestMessage
	}
	return errorResponse(http.StatusBadRequest, message)
}

// Unauthorized - Return a 401 Status Code with a message
func Unauthorized(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = unauthorizedMessage
	}
	return errorResponse(http.StatusUnauthorized, message)
}

// Forbidden - Return a 403 Status Code with a message
func Forbidden(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = forbiddenMessage
	}
	return errorResponse(http.StatusForbidden, message)
}

// NotFound - Return a 404 Status Code with a message
func NotFound(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = notFoundMessage
	}
	return errorResponse(http.StatusNotFound, message)
}

// MethodNotAllowed - Return a 405 Status Code with a message
func MethodNotAllowed(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = methodNotAllowedMessage
	}
	return errorResponse(http.StatusMethodNotAllowed, message)
}

// NotAcceptable - Return a 406 Status Code with a message
func NotAcceptable(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = notAcceptableMessage
	}
	return errorResponse(http.StatusNotAcceptable, message)
}

// PreconditionFailed - Return a 412 Status Code with a message
func PreconditionFailed(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = preconditionFailedMessage
	}
	return errorResponse(http.StatusPreconditionFailed, message)
}

// UnsupportedMediaType - Return a 415 Status Code with a message
func UnsupportedMediaType(message string) (events.ALBTargetGroupResponse, error) {
	if len(strings.TrimSpace(message)) == 0 {
		message = unsupportedMediaTypeMessage
	}
	return errorResponse(http.StatusUnsupportedMediaType, message)
}
