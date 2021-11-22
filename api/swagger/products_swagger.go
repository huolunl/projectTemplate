package docs

// swagger:route GET /v1/products/hello example someRequest
// some description.
// responses:
//   default: someResponse

// swagger:parameters someRequest
type someRequest struct {
	// This text will appear as description of your request body.
	// in:body
	// Required: true
	// some description
	Text string 
}

// swagger:response someResponse
type someResponse struct {
	// description code
	Code int 
	// description msg
	Message string 
	// description ref
	Reference string 
	// description data
	Data interface{}
}