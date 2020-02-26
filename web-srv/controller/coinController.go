package controller

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/kataras/iris"
	"github.com/sysdevguru/top-coin/web-srv/model"
	"github.com/sysdevguru/top-coin/web-srv/util"
)

// ListCoins returns list of top coins with limited amount
func ListCoins(ctx iris.Context) {
	// get limit and top
	limit := ctx.FormValue("limit")
	top := ctx.FormValue("top")
	// get type
	t := ctx.FormValue("type")

	// get coins and respond
	var coin model.Coin
	coins, err := coin.GetTopCoins(limit, top)
	if t == "csv" {
		header := []string{"rank", "symbol", "price"}
		b := &bytes.Buffer{}
		wr := csv.NewWriter(b)
		wr.Write(header)
		for _, v := range coins {
			price := fmt.Sprintf("%f", v.Price)
			record := []string{strconv.Itoa(v.Rank), v.Symbol, price}
			wr.Write(record)
		}
		wr.Flush()

		ctx.Header("Content-Type", "text/csv")
		w := ctx.ResponseWriter()
		w.Write(b.Bytes())
		return
	}
	if err != nil {
		util.SendResponse(ctx, iris.StatusInternalServerError, nil)
		return
	}
	util.SendResponse(ctx, iris.StatusOK, coins)
}
