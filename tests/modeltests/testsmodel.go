package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cepot-blip/fullstack/api/controllers"
	"github.com/cepot-blip/fullstack/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

//			DRIVER MENGGUNAKAN MYSQL
func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)
		if err != nil {
			fmt.Printf("Gagal connect to databases %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Success connected to databases %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Success reffresh table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		Nickname: "cepot",
		Email:    "cepotganteng@gmail.com",
		Password: "1qazxsw2",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("gagal membuat seed: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Nickname: "cepot ganteng",
			Email:    "cepot ganteng@gmail.com",
			Password: "1qazxsw2",
		},
		models.User{
			Nickname: "cepot ganteng lagi",
			Email:    "cepotganteng2@gmail.com",
			Password: "1qazxsw2",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}
	log.Printf("Success refresh table")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}
	user := models.User{
		Nickname: "cepot sans",
		Email:    "cepotsansbat@gmail.com",
		Password: "1qazxsw2",
	}
	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}
	post := models.Post{
		Title:    "Ini adalah title cepot",
		Content:  "langkah menjadi ganteng kaya cepot",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error

	if err != nil {
		return []models.User{}, []models.Post{}, err
	}
	var users = []models.User{
		models.User{
			Nickname: "cepot ganteng",
			Email:    "cepot ganteng@gmail.com",
			Password: "password",
		},
		models.User{
			Nickname: "cepot keren",
			Email:    "cepotkeren@gmail.com",
			Password: "1qazxsw2",
		},
	}
	var posts = []models.Post{
		models.Post{
			Title:   "Title 1 dari cepot",
			Content: "say hello to cepot",
		},
		models.Post{
			Title:   "Title 2 dari cepot lagi",
			Content: "say hello to cepot tamvan",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("gagal membuat seed pada table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("gagal membuat seed pada table: %v", err)
		}
	}
	return users, posts, nil
}
