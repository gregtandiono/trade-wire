package controller

import (
	"trade-wire/models"

	uuid "github.com/satori/go.uuid"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type CompanyController struct{}

func NewCompanyController() *CompanyController {
	return &CompanyController{}
}

func (cc *CompanyController) Save(ctx context.Context) {
	var company models.Company
	ctx.ReadJSON(&company)

	c := models.NewCompany(
		company.ID,
		company.Name,
		company.Address,
		company.CompanyType,
	)

	err := c.Save()

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not create company record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "company successfully created",
		})
	}

}

func (cc *CompanyController) FetchAll(ctx context.Context) {
	company := &models.Company{}
	companies := company.FetchAllCompanies()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(companies)
}

func (cc *CompanyController) FetchOne(ctx context.Context) {
	var company models.Company
	id := ctx.Params().Get("id")
	company.ID = uuid.FromStringOrNil(id)
	c := company.FetchOne()
	if c.ID == uuid.FromStringOrNil("") {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not find record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(c)
	}
}

func (cc *CompanyController) Update(ctx context.Context) {
	var company models.Company
	ctx.ReadJSON(&company)
	id := ctx.Params().Get("id")
	company.ID = uuid.FromStringOrNil(id)
	c := company.Update()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(c)
}

func (cc *CompanyController) Delete(ctx context.Context) {
	var company models.Company
	ctx.ReadJSON(&company)
	id := ctx.Params().Get("id")
	company.ID = uuid.FromStringOrNil(id)
	company.Delete()

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]string{
		"message": "company record successfully deleted",
	})
}
