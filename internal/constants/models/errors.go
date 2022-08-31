package models

type Errors struct {
	Err           error  `json:"error"`
	ErrorLocation string `json:"error_location"`
	ErrLine       int64  `json:"error_line"`
}
