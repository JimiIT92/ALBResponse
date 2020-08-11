/* Package albresponse - Package documentation
package albresponse
Functions:
	Response (status int, content interface{}, message string)
		Return a generic response with the specified Status Code and the specified content.
		The content could be any struct or primitive type, wich will then be returned
		into a JSON format. The message parameter will be used in case of error. If any error
		occurrs into the method, its value will be overwritten.

	Ok (content interface{})
		Return a 200 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	Create (content interface{})
		Return a 201 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	Accepted (content interface{})
		Return a 202 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	NoContent ()
		Return a 204 Status Code with no content response.

	MovedPermanently (content interface{})
		Return a 301 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	Found (content interface{})
		Return a 302 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	NotModified ()
		Return a 304 Status Code with no content response.

	TemporaryRedirect (content interface{})
		Return a 307 Status Code with the content response. The content
		could be any struct or primitive type, wich will then be returned
		into a JSON format.

	BadRequest (message string)
		Return a 400 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	Unauthorized (message string)
		Return a 401 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	Forbidden (message string)
		Return a 403 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	NotFound (message string)
		Return a 404 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	MethodNotAllowed (message string)
		Return a 405 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	NotAcceptable (message string)
		Return a 406 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	PreconditionFailed (message string)
		Return a 412 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	UnsupportedMediaType (message string)
		Return a 415 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	InternalServerError (message string)
		Return a 500 Status Code with a message. If the message is empty,
		than a default message will be displayed.

	NotImplemented (message string)
		Return a 501 Status Code with a message. If the message is empty,
		than a default message will be displayed.
*/

package albresponse // import "git-codecommit.eu-west-1.amazonaws.com/v1/repos/albresponse.git"
