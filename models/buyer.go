package models

import (
	"time"
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Buyer struct
type Buyer struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	PIC     string    `json:"pic"` // Person In Charge
}

// NewBuyer returns a new instance of Buyer struct
func NewBuyer(id uuid.UUID, name, address, pic string) *Buyer {
	return &Buyer{
		ID:      id,
		Name:    name,
		Address: address,
		PIC:     pic,
	}
}

// Save writes one buyer record to DB
func (b *Buyer) Save() error {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("buyers").Create(&Buyer{
		b.ID,
		b.Name,
		b.Address,
		b.PIC,
	})

	return nil
}

// FetchAllBuyers returns an array of buyers
func (b *Buyer) FetchAllBuyers() []Buyer {
	db := adaptors.DBConnector()
	defer db.Close()

	var buyers []Buyer
	db.Select([]string{"id", "name", "address", "pic"}).Where("deleted_at is null").Find(&buyers)
	return buyers
}

// FetchOne returns one buyer based on id in url param
func (b *Buyer) FetchOne() Buyer {
	db := adaptors.DBConnector()
	defer db.Close()

	var buyer Buyer
	db.Select([]string{"id", "name", "address", "pic"}).Where("id = ?", b.ID).Find(&buyer)
	return buyer
}

func (b *Buyer) Update() *Buyer {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("buyers").Where("id = ?", b.ID).Updates(&b)
	return b
}

func (b *Buyer) Delete() {
	db := adaptors.DBConnector()
	defer db.Close()

	db.Table("buyers").Where("id = ?", b.ID).Update("deleted_at", time.Now())
}
