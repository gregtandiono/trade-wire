package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Vessel model struct
type Vessel struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Beam   string    `json:"beam"`
	LOA    string    `json:"loa"`
	Draft  string    `json:"draft"`
	Status string    `json:"status"`
}

// NewVessel returns a new instance of Vessel struct
func NewVessel(id uuid.UUID, name, beam, loa, draft, status string) *Vessel {
	return &Vessel{
		ID:     id,
		Name:   name,
		Beam:   beam,
		LOA:    loa,
		Draft:  draft,
		Status: status,
	}
}

// Save creates a new vessel record in the db
func (v *Vessel) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	err := db.Table("vessels").Create(&Vessel{
		v.ID,
		v.Name,
		v.Beam,
		v.LOA,
		v.Draft,
		v.Status,
	}).Error

	return err
}

// FetchAllVessels returns an array of vessel records
func (v *Vessel) FetchAllVessels() ([]Vessel, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var vessels []Vessel
	err := db.Select([]string{"id", "name", "beam", "loa", "draft", "status"}).Where("deleted_at is null").Find(&vessels).Error
	return vessels, err
}

// FetchOne returns one vessel record based on id as search param
func (v *Vessel) FetchOne() (Vessel, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var vessel Vessel
	err := db.Select([]string{"id", "name", "beam", "loa", "draft", "status"}).Where("id = ?", v.ID).Find(&vessel).Error
	return vessel, err
}

// Update updates a vessel record in the db based on id
func (v *Vessel) Update() (*Vessel, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := v.FetchOne()

	if notFoundErr != nil {
		return v, notFoundErr
	}

	err := db.Table("vessels").Where("id = ?", v.ID).Updates(&v).Error
	return v, err
}

// Delete updates a records `deleted_at` column. Soft deletes
func (v *Vessel) Delete() error {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := v.FetchOne()
	if notFoundErr != nil {
		return notFoundErr
	}

	err := db.Table("commodities").Where("id = ?", v.ID).Update("deleted_at", time.Now()).Error
	return err
}
