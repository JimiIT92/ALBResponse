// Copyright 2019 Ecom Imaging
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
