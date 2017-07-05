package models

import (
	"time"
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

// FetchAllContacts returns an array of contacts
func (c *Contact) FetchAllContacts() ([]Contact, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var contacts []Contact
	err := db.Select([]string{"id", "name", "position", "office_number", "cell_number", "notes", "company_id"}).Where("deleted_at is null").Find(&contacts).Error
	return contacts, err
}

// FetchOne returns one record of contact based on record ID
func (c *Contact) FetchOne() (Contact, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var contact Contact
	err := db.Select([]string{"id", "name", "position", "office_number", "cell_number", "notes", "company_id"}).Where("id = ?", c.ID).Find(&contact).Error
	return contact, err
}

// Update updates an existing record
func (c *Contact) Update() (*Contact, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := c.FetchOne()

	if notFoundErr != nil {
		return c, notFoundErr
	}

	err := db.Table("contacts").Where("id = ?", c.ID).Updates(&c).Error
	return c, err
}

// Delete updates the `deleted_at` column in a record
func (c *Contact) Delete() error {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := c.FetchOne()
	if notFoundErr != nil {
		return notFoundErr
	}

	err := db.Table("contacts").Where("id = ?", c.ID).Update("deleted_at", time.Now()).Error
	return err
}
