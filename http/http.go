package global

// Import packages
import (
	"github.com/valyala/fasthttp"
)

// The HttpRequest struct contains five keys
/* - Client: *fasthttp.Client -> Client for sending http request 	*/
/* - Url: string -> The url to send http request to 				*/
/* - Method: string -> The method being used for http request 	*/
/* - Body: []byte -> The request body if request is put/post 		*/
/* - Headers: map[string]string -> The request headers map 		*/
type HttpRequest struct {
	Client  *fasthttp.Client
	Url     string
	Method  string
	Body    []byte
	Headers map[string]string
}

// The SetResponse() function is used to create a new fasthttp
// response that will be used for recieving the sent http request
// response data.
//
// The function requires the HttpRequest object which is used
// for determining whether to skip the request body by HEAD method
func SetResponse(req *HttpRequest) *fasthttp.Response {
	// Acquire the fasthttp response
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	// Whether to skip the response body
	response.SkipBody = req.Method == "HEAD"

	// Return the new response object
	return response
}

// The SetRequest() function is used to create a new fasthttp
// request that will be used for sending any http requests
//
// The function requires the HttpRequest object which is used
// for setting the request url, method, body and headers
func SetRequest(req *HttpRequest) *fasthttp.Request {
	// Acquire the fasthttp request
	var request *fasthttp.Request = fasthttp.AcquireRequest()

	// Set the Request Url and Method
	request.SetRequestURI(req.Url)
	request.Header.SetMethod(req.Method)

	// Set the request body
	if len(req.Body) > 0 {
		request.SetBody(req.Body)
	}

	// Set the request headers
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	// Return the new request object
	return request
}

// The Send() function is used to send an http request
// The function requires the HttpRequest object which contains
// the request url, method, client, headers and body
//
// The function releases the request object once the function
// closes
//
// The function returns the fasthttp response object and
// the request errors
func (_req *HttpRequest) Send() (*fasthttp.Response, error) {
	var (
		// Set the request object
		req *fasthttp.Request = SetRequest(_req)
		// Set the respone object
		resp *fasthttp.Response = SetResponse(_req)
		// Send the request and store any errors
		err error = _req.Client.Do(req, resp)
	)
	// Release the request once no longer needed
	defer fasthttp.ReleaseRequest(req)

	// Return the response and any client errors
	return resp, err
}
