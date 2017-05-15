package controller

import (
	"trade-wire/models"

	uuid "github.com/satori/go.uuid"

	iris "gopkg.in/kataras/iris.v6"
)

type CommodityController struct{}

func NewCommodityController() *CommodityController {
	return &CommodityController{}
}

func (cc *CommodityController) Save(ctx *iris.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)

	c := models.NewCommodity(
		commodity.ID,
		commodity.Name,
	)

	err := c.Save()

	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "failed to insert commodity record",
		})
	} else {
		ctx.JSON(iris.StatusOK, map[string]string{
			"message": "commodity successfully created",
		})
	}

}

func (cc *CommodityController) FetchAll(ctx *iris.Context) {
	commodity := &models.Commodity{}
	commodities, err := commodity.FetchAllCommodities()
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "failed to fetch all commodities",
		})
	} else {
		ctx.JSON(iris.StatusOK, commodities)
	}
}

func (cc *CommodityController) FetchOne(ctx *iris.Context) {
	var commodity models.Commodity
	id := ctx.Param("id")
	commodity.ID = uuid.FromStringOrNil(id)
	c, err := commodity.FetchOne()
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "could not find record",
		})
	} else {
		ctx.JSON(iris.StatusOK, c)
	}
}

func (cc *CommodityController) Update(ctx *iris.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)
	id := ctx.Param("id")
	commodity.ID = uuid.FromStringOrNil(id)
	c, err := commodity.Update()
	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "failed to update record",
		})
	} else {
		ctx.JSON(iris.StatusOK, c)
	}
}

func (cc *CommodityController) Delete(ctx *iris.Context) {
	var commodity models.Commodity
	ctx.ReadJSON(&commodity)
	id := ctx.Param("id")
	commodity.ID = uuid.FromStringOrNil(id)
	commodity.Delete()
	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "record successfully deleted",
	})
}
