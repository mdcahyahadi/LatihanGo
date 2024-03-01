package response

// Response model
type Response struct {
	ResponseCode string      `json:"rc"`
	Description  string      `json:"desc"`
	Message      string      `json:"msg"`
	Data         interface{} `json:"data"`
}

// ErrorStruct model dao
type ErrorStruct struct {
	Code        string `json:"code"`
	Description string `json:"desc,omitempty"`
	Message     string `json:"message"`
}

// we implement the built-in package 'error' interface by creating this function
func (e ErrorStruct) Error() string {
	return e.Message
}
