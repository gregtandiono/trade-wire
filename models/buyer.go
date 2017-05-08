package models

import (
	"trade-wire/adaptors"

	uuid "github.com/satori/go.uuid"
)

// Buyer struct
type Buyer struct {
	ID      uuid.UUID           `json:"id"`
	Name    string              `json:"name"`
	Address string              `json:"address"`
	PIC     []map[string]string `json:"pic"`
}

// NewBuyer returns a new instance of Buyer struct
func NewBuyer(id uuid.UUID, name, address string, pic []map[string]string) *Buyer {
	return &Buyer{
		ID:      id,
		Name:    name,
		Address: address,
		PIC:     pic,
	}
}

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
