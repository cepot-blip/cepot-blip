package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)


type Bank struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Bank_name	string		`gorm:"size:255;not null" json:"bank_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Bank)  Prepare(){
	b.ID				= 0
	b.Bank_name			= html.EscapeString(strings.TrimSpace(b.Bank_name))
	b.CreatedAt			= time.Now()
	b.UpdatedAt			= time.Now()
}


func (b *Bank)  Validate() error {

	if b.Bank_name == "" {
		return errors.New("requared bank name")
	}
	return nil
}

//		CREATE BANK
func (b *Bank) SaveBank(db *gorm.DB) (*Bank, error)  {
	var err error
	err = db.Debug().Model(&Bank{}).Create(&b).Error
	if err != nil {
		return &Bank{}, err
	}
	return b, nil
}

//		READ ALL BANK
func (b *Bank) FindAllBank(db *gorm.DB) (*[]Bank, error) {
	var err error
	banks := []Bank{}
	err = db.Debug().Model(&Bank{}).Limit(100).Find(&banks).Error
	if err != nil {
		return &[]Bank{}, err
	}
	return &banks, err
}


//		UPDATE BANK

func (b *Bank) UpdateBank(db *gorm.DB) (*Bank, error)  {
	var err error
	
	err = db.Debug().Model(&Bank{}).Where("id = ?", b.ID).Updates(Bank{Bank_name: b.Bank_name, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Bank{}, err
	}
	return b, nil
}


//		DELETE BANK

func (b *Bank) DeleteBank(db *gorm.DB, uid uint32) (int64, error) {
	
	db = db.Debug().Model(&Bank{}).Where("id = ?", uid).Take(&Bank{}).Delete(&Bank{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}