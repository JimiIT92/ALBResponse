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

// Package albresponse - Define functions for handling common 2xx Status Codes
package albresponse

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Ok - Return a 200 Status Code with the content response
func Ok(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusOK, content)
}

// Created - Return a 201 Status Code with the content response
func Created(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusCreated, content)
}

// Accepted - Return a 202 Status Code with the content response
func Accepted(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusAccepted, content)
}

// NoContent - Return a 204 Status Code
func NoContent() (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusNoContent, nil)
}
