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
	Bank		User		`json:"bank"`
	BankID		uint32		`gorm:"not null" json:"bank_id"`
	CreatedAt time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Bank)  Prepare(){
	b.ID				= 0
	b.Bank_name			= html.EscapeString(strings.TrimSpace(b.Bank_name))
	b.Bank				= User{}
	b.CreatedAt			= time.Now()
	b.UpdatedAt			= time.Now()
}


func (b *Bank)  Validate() error {

	if b.Bank_name == "" {
		return errors.New("requared bank name")
	}
	if b.BankID < 1 {
		return errors.New("required Bank")
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
	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", b.BankID).Take(&b.Bank).Error
		if err != nil {
			return &Bank{}, err
		}
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
	if len(banks) > 0 {
		for e := range banks {
			err := db.Debug().Model(&User{}).Where("id = ?", banks[e].BankID).Take(&banks[e].Bank).Error
			if err != nil {
				return &[]Bank{}, err
			}
		}
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
	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", b.BankID).Take(&b.Bank).Error
		if err != nil {
			return &Bank{}, err
		}
	}
	return b, nil
}


//		DELETE BANK

func (b *Bank) DeleteBank(db *gorm.DB, uid uint32, pid uint64) (int64, error) {
	
	db = db.Debug().Model(&Bank{}).Where("id = ? and bank_id = ?", pid, uid).Take(&Bank{}).Delete(&Bank{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Bank tidak ditemukan")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


//		FIND BANK BY ID

func (b *Bank) FindBankByID(db *gorm.DB, pid uint64) (*Bank, error) {
	var err error
	err = db.Debug().Model(&Bank{}).Where("id = ?", pid).Take(&b).Error
	if err != nil {
		return &Bank{}, err
	}
	if b.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", b.BankID).Take(&b.Bank).Error
		if err != nil {
			return &Bank{}, err
		}
	}
	return b, nil
}