package model_dify

// DifyRequest is the request for the Dify API
type DifyRequest struct {
	Message string `json:"message"`
}

// DifyResponse is the response from the Dify API
type DifyResponse struct {
	Title    string `json:"output_title"`
	Kind     string `json:"output_kind"`
	Contents string `json:"output_contents"`
}
