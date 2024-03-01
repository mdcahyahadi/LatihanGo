package uploadphoto

import "latihan/shared/response"

const (
	StatusSuccess string = "5"
	StatusFailed  string = "6"
)

const (
	FeatureTransferNow       string = "now"
	FeatureTransferScheduled string = "scheduled"
	FeatureTransferFX        string = "valas"
	FeatureTransferRecur     string = "recur"
)

const (
	FeatureIDOB     string = "100"
	FeatureIDIBT    string = "110"
	FeatureIDBIFast string = "180"
)

const (
	TopicNameTransferOB     string = "SCHEDULED-TRANSFER-OB"
	TopicNameTransferIBT    string = "SCHEDULED-TRANSFER-IBT"
	TopicNameTransferBIFast string = "SCHEDULED-TRANSFER-BI-FAST"
)

const (
	UserTypeMaker   int = 1
	UserTypeChecker int = 2
	UserTypeSigner  int = 3
	UserTypeAdmin   int = 4
)

var (
	ErrFeatureUnknown = response.ErrorStruct{Code: "PY01", Message: "unknown photo feature"}
	ErrUnAuthorized   = response.ErrorStruct{Code: "US06", Message: "unauthorized access"}
	ErrDataNotFound   = response.ErrorStruct{Code: "NF01", Message: "data not found"}
	ErrBadRequest     = response.ErrorStruct{Code: "BAD01", Message: "bad request"}
)

const (
	WaitingToProcessStatus int = 11
	CanceledStatus         int = 10
	FailedStatus           int = 6
	SuccessStatus          int = 5
	RejectStatus           int = 4
	ApprovedStatus         int = 3
	NeedApproveStatus      int = 2
	NeedVerificationStatus int = 1

	FailedStatusDescription  string = "FAILED"
	SuccessStatusDescription string = "SUCCESS"
)
