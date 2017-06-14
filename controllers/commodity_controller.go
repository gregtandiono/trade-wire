package controller

import (
	"trade-wire/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

type CommodityController struct{}

func NewCommodityController() *CommodityController {
	return &CommodityController{}
}

func (cc *CommodityController) Save(ctx context.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)

	c := models.NewCommodity(
		commodity.ID,
		commodity.Name,
	)

	err := c.Save()

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to insert commodity record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "commodity successfully created",
		})
	}

}

func (cc *CommodityController) FetchAll(ctx context.Context) {
	commodity := &models.Commodity{}
	commodities, err := commodity.FetchAllCommodities()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to fetch all commodities",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(commodities)
	}
}

func (cc *CommodityController) FetchOne(ctx context.Context) {
	var commodity models.Commodity
	id := ctx.Params().Get("id")
	commodity.ID = uuid.FromStringOrNil(id)
	c, err := commodity.FetchOne()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not find record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(c)
	}
}

func (cc *CommodityController) Update(ctx context.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)
	id := ctx.Params().Get("id")
	commodity.ID = uuid.FromStringOrNil(id)
	c, err := commodity.Update()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to update record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(c)
	}
}

func (cc *CommodityController) Delete(ctx context.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)
	id := ctx.Params().Get("id")
	commodity.ID = uuid.FromStringOrNil(id)
	err := commodity.Delete()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to delete record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "record successfully deleted",
		})
	}
}
