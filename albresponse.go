// Package albresponse - Define some functions for handling common response from an ALBTargetGroupRequest
package albresponse

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Return an error response
func errorResponse(status int, message string) (events.ALBTargetGroupResponse, error) {
	return Response(status, nil, message)
}

// Return a content response
func contentResponse(status int, content interface{}) (events.ALBTargetGroupResponse, error) {
	return Response(status, content, "")
}

// Response - Return a generic response
func Response(status int, content interface{}, message string) (events.ALBTargetGroupResponse, error) {
	body := ""
	var err error
	if content != nil {
		json, err := json.Marshal(content)
		if err != nil {
			status = http.StatusInternalServerError
		}
		body = string(json)
	}
	if err == nil && len(strings.TrimSpace(message)) > 0 {
		err = errors.New(message)
	}
	return events.ALBTargetGroupResponse{
		StatusCode:      status,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8;",
		},
		Body: body,
	}, err
}
