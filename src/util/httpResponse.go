package util

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type HttpErrors struct {
	Errors []string
}
