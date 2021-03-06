package plutus

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
)

// StaffKey is a struct that defines the strucutre of our staff_keys table. This table will hold the keys under our control that can be used in multisignature transactions.
type StaffKey struct {
	ID     uuid.UUID `json:"id" sql:"type:uuid" validate:"required,uuid" gorm:"primary"`
	PubKey []byte    `json:"public_key" sql:"type:text" validate:"required" gorm:"unique"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// GetStaffKey will search the database for all of our public keys, randomly select one, and return it to us.
func GetStaffKey() ([]byte, error) {
	var keys []*StaffKey
	err := Database.Find(&keys).Error
	if err != nil {
		return nil, err
	}

	// non cryptographically secure rand in use here, mainly because I don't really see the need, especially given that the set of pub keys will only ever get so large
	rand.Seed(time.Now().UTC().UnixNano())
	key := keys[rand.Intn(len(keys))].PubKey

	return key, nil
}

// Save will save this staff public key to our database
func (p *StaffKey) Save() error {
	err := p.validate()
	if err != nil {
		return err
	}
	return p.saveToDatabase()
}

func (p *StaffKey) saveToDatabase() error {
	return Database.Save(p).Error
}

func (p *StaffKey) validate() error {
	validate = validator.New()

	err := validate.Struct(p)
	if err != nil {
		return err
	}
	return nil
}
