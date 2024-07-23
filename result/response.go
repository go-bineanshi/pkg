package result

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int     `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Data    any     `json:"data,omitempty"`
	Errors  []error `json:"errors,omitempty"`
}

func NotAuthorized(ctx *gin.Context) {
	ctx.JSON(401, response{
		Code:    401,
		Message: "Unauthorized",
		Errors:  []error{ErrUnauthorized},
	})
}

func NotAllowed(ctx *gin.Context) {
	ctx.JSON(403, response{
		Code:    403,
		Message: "Forbidden",
		Errors:  []error{ErrForbidden},
	})
}

func Response(ctx *gin.Context, data any, errs ...error) {
	if len(errs) > 0 {
		ctx.JSON(200, response{
			Code:    400,
			Message: "业务错误",
			Errors:  errs,
		})
	} else {
		ctx.JSON(200, response{
			Code:    200,
			Message: "OK",
			Data:    data,
		})
	}
}
