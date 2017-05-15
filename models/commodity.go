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

func (c *Commodity) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	err := db.Table("commodities").Create(&Commodity{
		c.ID,
		c.Name,
	}).Error

	return err
}

func (c *Commodity) FetchAllCommodities() ([]Commodity, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var commodities []Commodity
	err := db.Select([]string{"id", "name"}).Where("deleted_at is null").Find(&commodities).Error
	return commodities, err
}

func (c *Commodity) FetchOne() (Commodity, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var commodity Commodity
	err := db.Select([]string{"id", "name"}).Where("id = ?", c.ID).Find(&commodity).Error
	return commodity, err
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
