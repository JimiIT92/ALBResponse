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

// Package albresponse - Define functions for handling common 3xx Status Codes
package albresponse

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// MovedPermanently - Return a 301 Status Code with the content response
func MovedPermanently(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusMovedPermanently, content)
}

// Found - Return a 302 Status Code with the content response
func Found(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusFound, content)
}

// SeeOther - Return a 302 Status Code with the content response
func SeeOther(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusSeeOther, content)
}

// NotModified - Return a 304 Status Code
func NotModified() (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusNotModified, nil)
}

// TemporaryRedirect - Return a 307 Status Code with the content response
func TemporaryRedirect(content interface{}) (events.ALBTargetGroupResponse, error) {
	return contentResponse(http.StatusTemporaryRedirect, content)
}
