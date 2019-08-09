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
