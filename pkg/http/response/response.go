package response

import (
	"github.com/adopabianko/dbo-service/pkg/stacktrace"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, httpCode int, message string, data, meta any) {
	res := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	ctx.JSON(httpCode, res)
}

func ResponseFailed(ctx *gin.Context, err error, message string) {
	code, appErr := stacktrace.Compile(err)

	res := ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   appErr,
	}

	ctx.AbortWithStatusJSON(int(code), res)
}

func BuildMeta(limit, totalItems, page, totalPages int, sortBy string) any {
	return gin.H{
		"items_per_page": limit,
		"total_items":    totalItems,
		"current_page":   page,
		"total_pages":    totalPages,
		"sort_by":        sortBy,
	}
}
