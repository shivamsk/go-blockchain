package response

type APIResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
}

type ErrorResponse struct {
	Data       interface{} `json:"errorMessage"`
	StatusCode int         `json:"statusCode"`
}
