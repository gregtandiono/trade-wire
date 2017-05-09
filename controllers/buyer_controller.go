package controller

import (
	"trade-wire/models"

	uuid "github.com/satori/go.uuid"

	iris "gopkg.in/kataras/iris.v6"
)

type BuyerController struct{}

func NewBuyerController() *BuyerController {
	return &BuyerController{}
}

func (bc *BuyerController) Save(ctx *iris.Context) {
	var buyer models.Buyer
	ctx.ReadJSON(&buyer)

	b := models.NewBuyer(
		buyer.ID,
		buyer.Name,
		buyer.Address,
		buyer.PIC,
	)

	err := b.Save()

	if err != nil {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "could not create buyer record",
		})
	} else {
		ctx.JSON(iris.StatusOK, map[string]string{
			"message": "buyer successfully created",
		})
	}

}

func (bc *BuyerController) FetchAll(ctx *iris.Context) {
	buyer := &models.Buyer{}
	buyers := buyer.FetchAllBuyers()
	ctx.JSON(iris.StatusOK, buyers)
}

func (bc *BuyerController) FetchOne(ctx *iris.Context) {
	var buyer models.Buyer
	id := ctx.Param("id")
	buyer.ID = uuid.FromStringOrNil(id)
	b := buyer.FetchOne()
	if b.ID == uuid.FromStringOrNil("") {
		ctx.JSON(iris.StatusBadRequest, map[string]string{
			"error": "could not find user",
		})
	} else {
		ctx.JSON(iris.StatusOK, b)
	}
}
