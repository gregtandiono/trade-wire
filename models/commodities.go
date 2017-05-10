package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

type Commodity struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewCommodity(id uuid.UUID, name string) *Commodity {
	return &Commodity{
		ID:   id,
		Name: name,
	}
}

// Save writes one buyer record to DB
func (c *Commodity) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("commodities").Create(&Commodity{
		c.ID,
		c.Name,
	})

	return nil
}

// FetchAllCommoditys returns an array of commodities
func (c *Commodity) FetchAllCommoditys() []Commodity {
	db := adaptors.DBConnector()
	defer db.Close()

	var commodities []Commodity
	db.Select([]string{"id", "name", "address", "pic"}).Where("deleted_at is null").Find(&commodities)
	return commodities
}

// FetchOne returns one buyer based on id in url param
func (c *Commodity) FetchOne() Commodity {
	db := adaptors.DBConnector()
	defer db.Close()

	var commodity Commodity
	db.Select([]string{"id", "name", "address", "pic"}).Where("id = ?", c.ID).Find(&commodity)
	return commodity
}
func (c *Commodity) Update() *Commodity {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("commodities").Where("id = ?", c.ID).Updates(&c)
	return c
}
func (c *Commodity) Delete() {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("commodities").Where("id = ?", c.ID).Update("deleted_at", time.Now())
}
