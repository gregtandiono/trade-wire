package models

import (
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Contact model struct
type Contact struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Position     string    `json:"position"`
	OfficeNumber string    `json:"office_number"`
	CellNumber   string    `json:"cell_number"`
	Notes        string    `json:"notes"`
	CompanyID    uuid.UUID `json:"company_id"`
}

// NewContact func returns a new instance of Contact struct
func NewContact(
	id, companyID uuid.UUID,
	name, position, officeNumber, cellNumber, notes string) *Contact {
	return &Contact{
		ID:           id,
		Name:         name,
		Position:     position,
		OfficeNumber: officeNumber,
		CellNumber:   cellNumber,
		Notes:        notes,
		CompanyID:    companyID,
	}
}

// Save func saves new contact record to db and returns an error (nullable)
func (c *Contact) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	err := db.Table("contacts").Create(&Contact{
		c.ID,
		c.Name,
		c.Position,
		c.OfficeNumber,
		c.CellNumber,
		c.Notes,
		c.CompanyID,
	}).Error

	return err
}
