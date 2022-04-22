package seed

import (
	"log"

	"github.com/cepot-blip/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "cepot ganteng",
		Email:    "cepotganteng@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "uta asu",
		Email:    "utaasu@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "tutorial muka glow up",
	},
	models.Post{
		Title:   "Title 2",
		Content: "tutorial muka glow down",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("gagal drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Admin{}).Error
	if err != nil {
		log.Fatalf("gagal migrasi table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("gagal membuat seed table users: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("gagal membuat seed table post: %v", err)
		}
	}

}
