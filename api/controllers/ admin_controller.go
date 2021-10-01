// package controllers

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/cepot-blip/fullstack/api/auth"
// 	"github.com/cepot-blip/fullstack/api/responses"
// 	"github.com/cepot-blip/fullstack/api/utils/formaterror"
// 	"github.com/cepot-blip/fullstack/models"
// 	"golang.org/x/crypto/bcrypt"
// )


// func (server *Server) LogIn(w http.ResponseWriter, r *http.Request) {
// 	body, err :=ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
	
// 	admin := models.Admin{}
// 	err = json.Unmarshal(body, &admin)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}

// 	admin.Prepare()
// 	err = admin.Validate("login")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	token, err := server.SignIn(admin.Email, admin.Password)
// 	if err != nil {
// 		formattedError := formaterror.FormatError(err.Error())
// 		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, token)
// }

// func (server *Server) Signin(email, password string) (string, error) {
// 	var err error
	
// 	admin := server.DB.Debug().Model(models.Admin{}).Where("email = ?", email).Take(&admin).Error
// 	if err != nil {
// 		return "", err
// 	}
// 	err = models.VerifyPassword(admin.Password, password)
// 	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
// 		return "", err
// 	}
// 	return auth.CreateToken(admin.ID)
// }