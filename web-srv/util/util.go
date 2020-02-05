package util

import (
	"github.com/kataras/iris"
)

func SendResponse(ctx iris.Context, statusCode int, data interface{}) {
	ctx.JSON(data)
	ctx.StatusCode(statusCode)
}
