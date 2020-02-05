package main

import (
	"log"
	"net"

	"github.com/kataras/iris"
	"github.com/sysdevguru/top-coin/web-srv/controller"
)

func main() {
	app := iris.New()

	// Just for CORS in case we have to implment admin server
	crs := func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")
		controller.AuthMdw(ctx)
	}

	// coin endpoints
	coinEndpoints := app.Party("/api/v1/coins", crs).AllowMethods(iris.MethodOptions)
	{
		coinEndpoints.Get("/list", controller.ListCoins)
	}

	// run listener
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Printf("Server listening Error:%v\n", err)
		return
	}
	app.Run(iris.Listener(listener))
}
