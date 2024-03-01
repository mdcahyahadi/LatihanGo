package response

import (
	"latihan/shared/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendSuccessResponse function with HTTPResponse
func SendSuccessResponse(c *gin.Context, resp Response) {
	c.JSON(http.StatusOK, resp)
}

// SendErrorResponse4xx function with HTTPResponse
func SendErrorResponse4xx(c *gin.Context, statusCode int, err ErrorStruct, data ...interface{}) {
	if utils.IsEmptyString(err.Description) {
		err.Description = DescriptionFailed
	}
	var resp Response
	if len(data) > 0 {
		resp = Response{
			ResponseCode: err.Code,
			Description:  err.Description,
			Message:      err.Message,
			Data:         data[0],
		}
	} else {
		resp = Response{
			ResponseCode: err.Code,
			Description:  err.Description,
			Message:      err.Message,
			Data:         nil,
		}

	}
	c.JSON(statusCode, resp)
}

// SendErrorResponse5xx function with HTTPResponse
func SendErrorResponse5xx(c *gin.Context, err ErrorStruct) {
	resp := Response{
		ResponseCode: err.Code,
		Description:  err.Description,
		Message:      err.Message,
		//Data:         nil,
	}

	c.JSON(http.StatusInternalServerError, resp)
}
