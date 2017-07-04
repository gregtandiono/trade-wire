package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Company struct
type Company struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	CompanyType string    `json:"company_type"`
}

// NewCompany takes in all key as param and returns a new instance of Company struct
func NewCompany(id uuid.UUID, name, address, companyType string) *Company {
	return &Company{
		ID:          id,
		Name:        name,
		Address:     address,
		CompanyType: companyType,
	}
}

func (c *Company) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("companies").Create(&Company{
		c.ID,
		c.Name,
		c.Address,
		c.CompanyType,
	})

	return nil
}

// FetchAllCompanies returns an array of Companies
func (c *Company) FetchAllCompanies() []Company {
	db := adaptors.DBConnector()
	defer db.Close()

	var Companies []Company
	db.Select([]string{"id", "name", "address", "company_type"}).Where("deleted_at is null").Find(&Companies)
	return Companies
}

// FetchOne returns one Company based on id in url param
func (c *Company) FetchOne() Company {
	db := adaptors.DBConnector()
	defer db.Close()

	var Company Company
	db.Select([]string{"id", "name", "address", "company_type"}).Where("id = ?", c.ID).Find(&Company)
	return Company
}

func (c *Company) Update() *Company {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("companies").Where("id = ?", c.ID).Updates(&c)
	return c
}

func (c *Company) Delete() {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("companies").Where("id = ?", c.ID).Update("deleted_at", time.Now())
}
