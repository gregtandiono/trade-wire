package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Variety model struct
type Variety struct {
	ID          uuid.UUID `json:"id"`
	CommodityID uuid.UUID `json:"commodity_id"`
	Name        string    `json:"name"`
	Origin      string    `json:"origin"`
	Specs       string    `json:"specs"`
}

// NewVariety func returns a new instance of Varietystruct
func NewVariety(
	id uuid.UUID,
	commodityID uuid.UUID,
	name string,
	origin string,
	specs string) *Variety {
	return &Variety{
		ID:          id,
		CommodityID: commodityID,
		Name:        name,
		Origin:      origin,
		Specs:       specs,
	}
}

func (v *Variety) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	err := db.Table("varieties").Create(&Variety{
		v.ID,
		v.CommodityID,
		v.Name,
		v.Origin,
		v.Specs,
	}).Error

	return err
}

func (v *Variety) FetchAllVarieties() ([]Variety, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var varieties []Variety
	err := db.Select([]string{"id", "name", "origin", "specs"}).Where("deleted_at is null").Find(&varieties).Error
	return varieties, err
}

func (v *Variety) FetchOne() (Variety, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var variety Variety
	err := db.Select([]string{"id", "name", "origin", "specs"}).Where("id = ?", v.ID).Find(&variety).Error
	return variety, err
}
func (v *Variety) Update() (*Variety, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := v.FetchOne()

	if notFoundErr != nil {
		return v, notFoundErr
	}

	err := db.Table("varieties").Where("id = ?", v.ID).Updates(&v).Error
	return v, err
}

func (v *Variety) Delete() error {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := v.FetchOne()
	if notFoundErr != nil {
		return notFoundErr
	}

	err := db.Table("varieties").Where("id = ?", v.ID).Update("deleted_at", time.Now()).Error
	return err
}
