package global

// Import packages
import (
	"github.com/valyala/fasthttp"
)

// HttpRequest struct to hold headers, method, url, etc.
type HttpRequest struct {
	Client  *fasthttp.Client
	Url     string
	Method  string
	Headers map[string]string
}

// Function to return a fasthttp response
func SetResponse(req *HttpRequest) *fasthttp.Response {
	// Acquire the fasthttp response
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	// Whether to skip the response body
	response.SkipBody = req.Method == "HEAD"

	// Return the new response object
	return response
}

// Function to return a fasthttp request
func SetRequest(req *HttpRequest) *fasthttp.Request {
	// Acquire the fasthttp request
	var request *fasthttp.Request = fasthttp.AcquireRequest()

	// Set the Request Url and Method
	request.SetRequestURI(req.Url)
	request.Header.SetMethod(req.Method)

	// Set the request headers
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}

	// Return the new request object
	return request
}

// Function to send an http request using the HttpRequest struct
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
