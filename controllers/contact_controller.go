package controller

import (
	"trade-wire/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	uuid "github.com/satori/go.uuid"
)

// ContactController empty struct
type ContactController struct{}

// NewContactController returns a new instance of ContactController struct
func NewContactController() *ContactController {
	return &ContactController{}
}

// Save initiates save model function
func (cc *ContactController) Save(ctx context.Context) {
	var contact models.Contact
	ctx.ReadJSON(&contact)

	c := models.NewContact(
		contact.ID,
		contact.CompanyID,
		contact.Name,
		contact.Position,
		contact.OfficeNumber,
		contact.CellNumber,
		contact.Notes,
	)

	err := c.Save()

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(map[string]string{
			"error": "could not create contact record",
		})
	} else {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(map[string]string{
			"message": "contact successfully created",
		})
	}
}

// FetchAll initiates FetchAllContacts model func
func (cc *ContactController) FetchAll(ctx context.Context) {
	contact := &models.Contact{}
	contacts := contact.FetchAllContacts()
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(contacts)
}

// FetchOne initiates FetchOne model func
func (cc *ContactController) FetchOne(ctx context.Context) {
	var contact models.Contact
	id := ctx.Params().Get("id")
	contact.ID = uuid.FromStringOrNil(id)
	c, err := contact.FetchOne()
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

// Update invokes the Update model func
func (cc *ContactController) Update(ctx context.Context) {
	var contact models.Contact
	ctx.ReadJSON(&contact)
	id := ctx.Params().Get("id")
	contact.ID = uuid.FromStringOrNil(id)
	c, err := contact.Update()
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

// Delete invokes the Delete model func
func (cc *ContactController) Delete(ctx context.Context) {
	var contact models.Contact
	ctx.ReadJSON(&contact)
	id := ctx.Params().Get("id")
	contact.ID = uuid.FromStringOrNil(id)
	err := contact.Delete()

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
