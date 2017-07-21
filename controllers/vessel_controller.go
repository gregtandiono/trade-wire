package controller

import (
	"trade-wire/adaptors"
	"trade-wire/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

// VesselController empty struct
type VesselController struct{}

// NewVesselController returns an instance of VesselController empty struct
func NewVesselController() *VesselController {
	return &VesselController{}
}

// Save triggers the {model.save} model method
func (vc *VesselController) Save(ctx context.Context) {
	var vessel models.Vessel
	ctx.ReadJSON(&vessel)

	v := models.NewVessel(
		vessel.ID,
		vessel.Name,
		vessel.Beam,
		vessel.LOA,
		vessel.Draft,
		vessel.Status,
	)

	err := v.Save()
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(adaptors.ResponseTemplate("insert:fail"))
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(adaptors.ResponseTemplate("insert:success"))
	}
}

func (vc *VesselController) FetchAll(ctx context.Context) {
	var vessel models.Vessel
	vessels, err := vessel.FetchAllVessels()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("fetch:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(vessels)
	}
}

func (vc *VesselController) FetchOne(ctx context.Context) {
	var vessel models.Vessel
	id := ctx.Params().Get("id")
	vessel.ID = uuid.FromStringOrNil(id)
	v, err := vessel.FetchOne()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("fetch:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(v)
	}
}

func (vc *VesselController) Update(ctx context.Context) {
	var vessel models.Vessel
	id := ctx.Params().Get("id")
	vessel.ID = uuid.FromStringOrNil(id)
	v, err := vessel.Update()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("update:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(v)
	}
}

func (vc *VesselController) Delete(ctx context.Context) {
	var vessel models.Vessel
	id := ctx.Params().Get("id")
	vessel.ID = uuid.FromStringOrNil(id)
	err := vessel.Delete()
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(adaptors.ResponseTemplate("delete:fail"))
	} else {
		ctx.StatusCode(200)
		ctx.JSON(adaptors.ResponseTemplate("delete:success"))
	}
}
