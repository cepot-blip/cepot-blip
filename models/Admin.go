package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)


type Admin struct {
	ID			uint32		`gorm:"primary_key;auto_increment" json:"id"`
	Email		string		`gorm:"size:100;not null;unique" json:"email"`
	Password	string		`gorm:"size:255;not null;" json:"password"`
	CreatedAt time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}


//	HASH PASSWORD
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//	COMPARE PASSWORD
func VerifyPassword1(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (a *Admin) BeforeSave()error{
	hashedPassword, err := Hash(a.Password)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}

func (a *Admin) Prepare() {
	a.ID = 0
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

//		VALIDASI

func (a *Admin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if a.Email == "" {
			return errors.New("required email anda")
		}
		if a.Password == "" {
			return errors.New("required password anda")
		}
		if err := checkmail.ValidateFormat(a.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	case "login":
		if a.Password == "" {
			return errors.New("require password")
		}
		if a.Email == "" {
			return errors.New("require email")
		}
		if err := checkmail.ValidateFormat(a.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if a.Email == "" {
			return errors.New("require email")
		}
		if a.Password == "" {
			return errors.New("require Password")
		}
		if err := checkmail.ValidateFormat(a.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil
	}
}


//		CREATE ADMIN
func (a *Admin) SaveAdmin(db *gorm.DB) (*Admin, error) {
	var _,err  error
	err = db.Debug().Create(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}

//		READ ALL ADMIN

func (a *Admin) FindAllAdmin(db *gorm.DB)(*[]Admin, error) {
	var err	error
	admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Limit(100).Find(&admins).Error
	if err != nil {
		return &[]Admin{}, err
	}
	return &admins, err
}

//		LOGIN ADMIN BY ID

func (a *Admin)FindAdminByID(db *gorm.DB, uid uint32) (*Admin, error)  {
	var _,err error
	err = db.Debug().Model(Admin{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &Admin{},err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("Admin tidak ditemukan")
	}
	return a, err
}

//	 	UPDATE ADMIN
func (a *Admin) UpdateAdmin(db *gorm.DB, uid uint32) (*Admin, error)  {
	err := a.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"email" 	: a.Email,
			"password" 	: a.Password,
			"update_at" : time.Now(),
		},
	)
	if db.Error != nil {
		return &Admin{}, db.Error
	}

	//		tampilan admin setelah diupdate
	err = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&a).Error
	if err != nil {
		return &Admin{}, err
	}
	return a, nil
}


//		DELETE ADMINS
func (a *Admin) DeleteAdmin(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).Delete(&Admin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}