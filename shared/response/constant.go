package response

const (
	RCSuccess      string = "00"
	RCErrorGeneral string = "99"
)

const (
	DescriptionSuccess string = "success"
	DescriptionFailed  string = "failed"
	ErrorGeneral       string = "general error"
)

var (
	ErrOverride = ErrorStruct{
		Code:    "",
		Message: "",
	}

	ErrBadRequest = ErrorStruct{Code: "BAD01", Message: "Bad Request"}
	ErrNotFound   = ErrorStruct{Code: "NTF01", Message: "Data Not Found"}
)

// APM Const
const (
	APMSpanTypeHandlerProcess string = "handler_process"
	APMSpanTypeLogicalProcess string = "logical_process"

	APMTrxCustomKeyRequestParam      string = "request_param"
	APMTrxCustomKeyRequestQueryParam string = "request_query_param"
	APMTrxCustomKeyRequestBody       string = "request_body"
	APMTrxCustomKeyResponse          string = "response"
)
