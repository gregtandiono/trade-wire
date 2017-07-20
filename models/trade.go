package models

import (
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Trade model struct
type Trade struct {
	ID         uuid.UUID `json:"id"`
	CompanyID  uuid.UUID `json:"company_id"`
	VarietyID  uuid.UUID `json:"variety_id"`
	VesselID   uuid.UUID `json:"vessel_id"`
	Quantity   int
	BLQuantity int
	Shipment   string
	Price      int
	PriceNote  string
	Status     string
	Note       string
}

// NewTrade returns a new instance of Trade struct
func NewTrade(
	id, companyID, varietyID, vesselID uuid.UUID,
	quantity, blQuantity, price int,
	shipment, priceNote, status, note string) *Trade {
	return &Trade{
		ID:         id,
		CompanyID:  companyID,
		VarietyID:  varietyID,
		VesselID:   vesselID,
		Quantity:   quantity,
		BLQuantity: blQuantity,
		Shipment:   shipment,
		Price:      price,
		PriceNote:  priceNote,
		Status:     status,
		Note:       note,
	}
}

// Save creates a new Trade record in the db
func (t *Trade) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	err := db.Table("trades").Create(&Trade{
		t.ID,
		t.CompanyID,
		t.VarietyID,
		t.VesselID,
		t.Quantity,
		t.BLQuantity,
		t.Shipment,
		t.Price,
		t.PriceNote,
		t.Status,
		t.Note,
	}).Error

	return err
}

//
// // FetchAllVessels returns an array of vessel records
// func (v *Vessel) FetchAllVessels() ([]Vessel, error) {
// 	db := adaptors.DBConnector()
// 	defer db.Close()
//
// 	var vessels []Vessel
// 	err := db.Select([]string{"id", "name", "beam", "loa", "draft", "status"}).Where("deleted_at is null").Find(&vessels).Error
// 	return vessels, err
// }
//
// // FetchOne returns one vessel record based on id as search param
// func (v *Vessel) FetchOne() (Vessel, error) {
// 	db := adaptors.DBConnector()
// 	defer db.Close()
//
// 	var vessel Vessel
// 	err := db.Select([]string{"id", "name", "beam", "loa", "draft", "status"}).Where("id = ?", v.ID).Find(&vessel).Error
// 	return vessel, err
// }
//
// // Update updates a vessel record in the db based on id
// func (v *Vessel) Update() (*Vessel, error) {
// 	db := adaptors.DBConnector()
// 	defer db.Close()
//
// 	_, notFoundErr := v.FetchOne()
//
// 	if notFoundErr != nil {
// 		return v, notFoundErr
// 	}
//
// 	err := db.Table("vessels").Where("id = ?", v.ID).Updates(&v).Error
// 	return v, err
// }
//
// // Delete updates a records `deleted_at` column. Soft deletes
// func (v *Vessel) Delete() error {
// 	db := adaptors.DBConnector()
// 	defer db.Close()
//
// 	_, notFoundErr := v.FetchOne()
// 	if notFoundErr != nil {
// 		return notFoundErr
// 	}
//
// 	err := db.Table("commodities").Where("id = ?", v.ID).Update("deleted_at", time.Now()).Error
// 	return err
// }
