package controller

import (
	"github.com/kataras/iris"
)

// AuthMdw handles authentication
func AuthMdw(ctx iris.Context) {
	// get token and check token

	ctx.Next()
}
