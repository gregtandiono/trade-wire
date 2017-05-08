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
	var buyer *models.Buyer
	ctx.ReadJSON(&buyer)
	buyer.ID = uuid.NewV4()

	// fmt.Println(buyer)
	// models.Save("buyers", &buyer)
	b := models.NewAbstractRecord{
		uuid.NewV4(),
		buyer.Name,
		buyer.Address,
		buyer.PIC,
	}

	b.Save("buyers")

	ctx.JSON(iris.StatusOK, map[string]string{
		"message": "buyer successfully created",
	})
}
