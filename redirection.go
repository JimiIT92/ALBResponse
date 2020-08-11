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
