package controller

import (
	"trade-wire/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

type VarietyController struct{}

func (vc *VarietyController) Save(ctx context.Context) {
	var variety models.Variety
	ctx.ReadJSON(&variety)

	v := models.NewVariety(
		variety.ID,
		variety.CommodityID,
		variety.Name,
		variety.Origin,
		variety.Specs,
	)

	err := v.Save()

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to insert variety record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "variety successfully created",
		})
	}
}

func (vc *VarietyController) FetchAll(ctx context.Context) {
	variety := &models.Variety{}
	varieties, err := variety.FetchAllVarieties()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "failed to fetch all varieties",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(varieties)
	}
}

func (vc *VarietyController) FetchOne(ctx context.Context) {
	var variety models.Variety
	id := ctx.Params().Get("id")
	variety.ID = uuid.FromStringOrNil(id)
	v, err := variety.FetchOne()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not find record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(v)
	}
}

func (vc *VarietyController) Update(ctx context.Context) {
	var variety models.Variety
	ctx.ReadJSON(&variety)
	id := ctx.Params().Get("id")
	variety.ID = uuid.FromStringOrNil(id)
	c, err := variety.Update()
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

func (vc *VarietyController) Delete(ctx context.Context) {
	var variety models.Variety
	ctx.ReadJSON(&variety)
	id := ctx.Params().Get("id")
	variety.ID = uuid.FromStringOrNil(id)
	err := variety.Delete()
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
