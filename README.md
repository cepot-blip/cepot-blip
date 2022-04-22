# FULL STACK REST API IN GOLANG PROFESIONAL DEVELOPER

---
![go](https://user-images.githubusercontent.com/85933775/164717883-0511997e-a405-4780-8d82-d310ed6996ae.gif)

**LANGKAH PERTAMA UNTUK MEMBUAT REST FULL API DENGAN GOLANG YAITU,**

1. Membuat folder project bernama **Fullstack**
2. Inisialisai Poroject dengan cara melakukan syntax berikut .

```go

go mod init github.com/{username git}/{nama folder}

// contoh
go mod init github.com/cepot-blip/fullstack

```

3. Instalasi beberapa dependensi kebutuahn project kita, disini saya menggunakan **ORM gorm**

```go
go get github.com/jinzhu/gorm
```

4. Install bcrypt dengan melakukan syntax berikut, untuk melakukan penghashan pada password
5. Install Gorila mux untuk Router nya syntax sebagai berikut

```go
go get github.com/gorilla/mux
```

```go
go get golang.org/x/crypto/bcrypt
```

6. Install JWT atau Json Web Token untuk keperluan login pada user atau lainnya sbt .

```go
go get github.com/dgrijalva/jwt-go

```

7. Install Databases, karena disini saya menggunakan Databases Mysql untuk keperluan Project saya dan sebagai berikut syntax ny

```go
go get github.com/jinzhu/gorm/dialects/mysql"
```

**Setalah selesai semua kebutan dependensi yang kita butuhkan langsung masuk kedalam Kode editor kesayang kita semua Visual Studio Code, dan langkah selanjutanya yaitu menentukan folder atau bisa di sebut dengan FOLDER STRUKTURING agar mempermudahkan dalam pengerjaan project kita**

## Didalam folder project Fullstack yang tadi kita buat masuk kedalam nya dan kita akan membuat Folder dan File bernama API ,TEST, .ENV dan main.go

- api

_folder dibawah ini ada di dalam semua folder api_

- auth
- controllers
- middlewares
- models
- responses
- seed
- utils

\*Dan didalam folder api membuat file bernama **server.go\***
\*Dan DILUAR folder api membuat file bernama **main.go untuk runing server utama kita\***

_Dan tampilan nya akan seperti ini jika sudah membuat folder semua itu._


![sss](https://user-images.githubusercontent.com/85933775/164716908-399f4c1e-71b0-455d-ba9a-ef383bc93721.PNG)



> **selanjut nya kita buka file **.env** kita untuk mengedit isi dari file tersebut dengan membuat seperti berikut.**
> 
![env](https://user-images.githubusercontent.com/85933775/164713812-9541f888-0321-42de-bfe8-5a1ad1a961f7.PNG)


Berhubunga kita menggunakan mysql sebagai DB nya kita membutuhkan DB_USER mysql dan DB_PASSWORD kita dan untuk DB_NAME itu kita membuat nya di dalam databases mysql kita bernama server_golang sebagai contoh, dan utuk db_port karena saya menggunakan window jadi port nya **3306** dan untuk pengguna mac os port nya **8889, jangan lupa untuk start mamp agar terconect ke server mysql nya, silahkan di download apk mamp nya di google banyak.**

> SELANJUT NYA **kita akan membuat file di dalam folder _controllers_ bernama _base.go dimana sebagai controller untuk file server.go kita_**

```go
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 9000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT TOKEN FILE DIDALAM FOLDER AUTH bernama token_.go untuk Membuat token pada users untuk validasi login**

```go
package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

//		CREATE TOKEN USERS
func CreateToken(user_id uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expayed sebelum 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

//		EXTRACT TOKEN
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

//		EXTRACT TOKEN ID
func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id, admins_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

// 	Cukup tampilkan klaim licely di terminal
func Pretty(data interface{}) {
	e, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(e))
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE DIDALAM FOLDER RESPONSES bernama handlejson_.go untuk Menjadikan json,**

```go
package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FOLDER DIDALAM FOLDER UTILS bernama FORMATERROR didalam FOLDER FORMATERROR buat file bernama handleerror_.go untuk Menghandle apa bila terjadi error,**

```go
package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "nickname") {
		return errors.New("Nickname must be unique")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email must be unique")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE DIDALAM FOLDER MODEL bernama Users_.go untuk membuat Model Users Apa saja yang di butuhkan,**

```go
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

//		MODEL YANG INGIN DI BUAT
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//		HASH PASSWORD
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//		COMPARE PASSWORD
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

//		VALIDASI
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("required Nickname")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("required Nickname")
		}
		if u.Password == "" {
			return errors.New("required Password")
		}
		if u.Email == "" {
			return errors.New("required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid Email")
		}
		return nil
	}
}

//		CREATE USERS
func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var _, err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//		READ ALL USERS
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

//		LOGIN USERS BY ID
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var _, err error
	err = db.Debug().Model([]User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

//		UPDATE USERS
func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// Untuk hash password kembali
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	//   Ini adalah tampilan users yang diperbarui
	err = db.Debug().Model(&User{}).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//		DELETE USERS
func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE DIDALAM FOLDER CONTROLLERS bernama users_controllers_.go untuk Mengcontroller Model Users yang sudah kita buat tadi,**

```go
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cepot-blip/fullstack/api/auth"
	"github.com/cepot-blip/fullstack/api/models"
	"github.com/cepot-blip/fullstack/api/responses"
	"github.com/cepot-blip/fullstack/api/utils/formaterror"
	"github.com/gorilla/mux"
)

//		CREATE USERS
func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusOK, userCreated)
}

//		READ ALL USERS
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

//		FIND USERS BY ID
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGet, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGet)
}

//		UPDATE USERS
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars[""], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	 if err != nil {
	 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	 	return
	 }
	 if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
 	return
	 }
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
}

//		DELETE USERS
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	 tokenID, err := auth.ExtractTokenID(r)
	 if err != nil {
	 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	 	return
	 }
	 if tokenID != 0 && tokenID != uint32(uid) {
	 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	 	return
	 }
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE DI DALAM FOLDER SEED bernama_ seed.go untuk mengkonekan file server.go agar terkonek ke DB atau untuk ngeload data dari databases ketika sudah dibuatkan model**

```go
package seed

import (
	"log"

	"github.com/cepot-blip/fullstack/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "pak tarno rasa leci",
		Email:    "paktarno@gmail.com",
		Password: "password",
	},
	models.User{
		Nickname: "udin rasa kecap",
		Email:    "udinrasakecap@gmail.com",
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

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("gagal drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("gagal migrasi table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("gagal membuat seed table users: %v", err)
		}
	}
}
```

> **_SELANJUT NYA KITA AKAN MENGEDIT FILE_** **server.go untuk mengkonekan file .env kita ke server.go**

```go
package api

import (
	"fmt"
	"log"
	"os"

	"github.com/cepot-blip/fullstack/api/controllers"
	"github.com/cepot-blip/fullstack/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error dalam mendapatkan file env, not comming through %v", err)
	} else {
		fmt.Println("berhasil mendapatkan env")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":9000")

}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE_ middlewares.go di dalam folder MIDDLEWARES untuk kepetingan routing pada END POINT,**

```go
package middlewares

import (
	"errors"
	"net/http"

	"github.com/cepot-blip/fullstack/api/auth"
	"github.com/cepot-blip/fullstack/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := (r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
```

> **_SELANJUT NYA KITA AKAN MEMBUAT FILE_** **routes.go untuk mebuat EndPoint dan memanggil model yang sudah kita buat ,**

```go
package controllers

import "github.com/cepot-blip/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {

	// 		Home Routes
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

			// Login Routes USER
	s.Router.HandleFunc("/user_login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//		Users routess
	s.Router.HandleFunc("/users_create", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users_read", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users_update", middlewares.SetMiddlewareJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users_delete/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser)).Methods("DELETE")
}
```

> **_SELANJUT LANGKAH TERAKHIR KITA AKAN MENGEDIT FILE_** main**.go untuk memanggil folder api kita untuk meruning struktur foldeer yang sudah kita buat,**

```go
package main

import "github.com/cepot-blip/fullstack/api"

func main() {
	api.Run()
}
```

# DAN UNTUK MERUNING SERVER KITA DI TERMINAL SILAHKAN KETIK SYNTAX BERIKUT,

- **_nodemon --exec go run _.go -signal SIGTERM\***

**atau bisa juga dengan menjalankan file utama kita yatu**

- **go run main.go**
