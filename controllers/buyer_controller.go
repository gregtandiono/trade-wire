package controller

import (
	"trade-wire/models"

	uuid "github.com/satori/go.uuid"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type BuyerController struct{}

func NewBuyerController() *BuyerController {
	return &BuyerController{}
}

func (bc *BuyerController) Save(ctx context.Context) {
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
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not create buyer record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "buyer successfully created",
		})
	}

}

func (bc *BuyerController) FetchAll(ctx context.Context) {
	buyer := &models.Buyer{}
	buyers := buyer.FetchAllBuyers()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(buyers)
}

func (bc *BuyerController) FetchOne(ctx context.Context) {
	var buyer models.Buyer
	id := ctx.Params().Get("id")
	buyer.ID = uuid.FromStringOrNil(id)
	b := buyer.FetchOne()
	if b.ID == uuid.FromStringOrNil("") {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not find record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(b)
	}
}

func (bc *BuyerController) Update(ctx context.Context) {
	var buyer models.Buyer
	ctx.ReadJSON(&buyer)
	id := ctx.Params().Get("id")
	buyer.ID = uuid.FromStringOrNil(id)
	b := buyer.Update()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(b)
}

func (bc *BuyerController) Delete(ctx context.Context) {
	var buyer models.Buyer
	ctx.ReadJSON(&buyer)
	id := ctx.Params().Get("id")
	buyer.ID = uuid.FromStringOrNil(id)
	buyer.Delete()

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]string{
		"message": "buyer record successfully deleted",
	})
}
