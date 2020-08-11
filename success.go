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
