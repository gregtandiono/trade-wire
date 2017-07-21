package controller

import (
	"trade-wire/adaptors"
	"trade-wire/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

// TradeController empty struct
type TradeController struct{}

// NewTradeController returns an instance of TradeController empty struct
func NewTradeController() *TradeController {
	return &TradeController{}
}

//
// Save triggers the {model.save} model method
func (tc *TradeController) Save(ctx context.Context) {
	var trade models.Trade
	ctx.ReadJSON(&trade)

	t := models.NewTrade(
		trade.ID,
		trade.BuyerID,
		trade.SupplierID,
		trade.VarietyID,
		trade.VesselID,
		trade.Quantity,
		trade.BLQuantity,
		trade.Price,
		trade.Shipment,
		trade.PriceNote,
		trade.Status,
		trade.Notes,
	)

	err := t.Save()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(adaptors.ResponseTemplate("insert:fail"))
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(adaptors.ResponseTemplate("insert:success"))
	}
}

// FetchAll returns an array of trade objects
func (tc *TradeController) FetchAll(ctx context.Context) {
	var trade models.Trade
	trades, err := trade.FetchAllTrades()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("fetch:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(trades)
	}
}

// FetchOne returns one trade object based on ID param in req body
func (tc *TradeController) FetchOne(ctx context.Context) {
	var trade models.Trade
	id := ctx.Params().Get("id")
	trade.ID = uuid.FromStringOrNil(id)
	t, err := trade.FetchOne()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("fetch:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(t)
	}
}

// Update updates a trade record from the DB based on the ID param in the req body
func (tc *TradeController) Update(ctx context.Context) {
	var trade models.Trade
	id := ctx.Params().Get("id")
	trade.ID = uuid.FromStringOrNil(id)
	t, err := trade.Update()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("update:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(t)
	}
}

// Delete soft deletes a record by updating the `deleted_at` column of the record
func (tc *TradeController) Delete(ctx context.Context) {
	var trade models.Trade
	id := ctx.Params().Get("id")
	trade.ID = uuid.FromStringOrNil(id)
	err := trade.Delete()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("delete:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(adaptors.ResponseTemplate("delete:success"))
	}
}
