package controller

import (
	"github.com/kataras/iris"
	"github.com/sysdevguru/top-coin/web-srv/model"
	"github.com/sysdevguru/top-coin/web-srv/util"
)

// ListCoins returns list of top coins with limited amount
func ListCoins(ctx iris.Context) {
	// get limit
	limit := ctx.FormValue("limit")

	// get coins and respond
	var coin model.Coin
	coins, err := coin.GetTopCoins(limit)
	if err != nil {
		util.SendResponse(ctx, iris.StatusInternalServerError, nil)
		return
	}
	util.SendResponse(ctx, iris.StatusOK, coins)
}
