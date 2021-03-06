package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Trade model struct
type Trade struct {
	ID         uuid.UUID `json:"id"`
	CompanyID  uuid.UUID `json:"company_id"`
	VarietyID  uuid.UUID `json:"variety_id"`
	VesselID   uuid.UUID `json:"vessel_id"`
	Quantity   int       `json:"quantity"`
	BLQuantity int       `json:"bl_quantity"`
	Shipment   string    `json:"shipment"`
	Price      int       `json:"price"`
	PriceNote  string    `json:"price_note"`
	Status     string    `json:"status"`
	Note       string    `json:"note"`
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

// FetchAllTrades returns an array of vessel records
func (t *Trade) FetchAllTrades() ([]Trade, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var trades []Trade
	err := db.Select([]string{
		"id", "company_id", "variety_id",
		"vessel_id", "quantity", "bl_quantity", "shipment", "price",
		"price_note", "status", "note"}).Where("deleted_at is null").Find(&trades).Error
	return trades, err
}

// FetchOne returns one trade record based on id as search param
func (t *Trade) FetchOne() (Trade, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	var trade Trade
	err := db.Select([]string{
		"id", "company_id", "variety_id",
		"vessel_id", "quantity", "bl_quantity", "shipment", "price",
		"price_note", "status", "note"}).Where("id = ?", t.ID).Find(&trade).Error
	return trade, err
}

// Update updates a vessel record in the db based on id
func (t *Trade) Update() (*Trade, error) {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := t.FetchOne()

	if notFoundErr != nil {
		return t, notFoundErr
	}

	err := db.Table("trades").Where("id = ?", t.ID).Updates(&t).Error
	return t, err
}

// Delete updates a record's `deleted_at` column. Soft deletes
func (t *Trade) Delete() error {
	db := adaptors.DBConnector()
	defer db.Close()

	_, notFoundErr := t.FetchOne()
	if notFoundErr != nil {
		return notFoundErr
	}

	err := db.Table("trades").Where("id = ?", t.ID).Update("deleted_at", time.Now()).Error
	return err
}
