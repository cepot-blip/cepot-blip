package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cepot-blip/fullstack/api/auth"
	"github.com/cepot-blip/fullstack/api/responses"
	"github.com/cepot-blip/fullstack/api/utils/formaterror"
	"github.com/cepot-blip/fullstack/models"
	"golang.org/x/crypto/bcrypt"
)

//		USER LOGIN
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
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

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}


//		ADMIN LOGIN
func (server *Server) LoginAdmin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	admins := models.Admin{}
	err = json.Unmarshal(body, &admins)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	admins.Prepare()
	err = admins.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(admins.Email, admins.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignInAdmin(email, password string) (string, error) {

	var err error

	admins := models.Admin{}

	err = server.DB.Debug().Model(models.Admin{}).Where("email = ?", email).Take(&admins).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(admins.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(admins.ID)
	admin := models.Admin{}

	err = server.DB.Debug().Model(models.Admin{}).Where("email = ?", email).Take(&admin).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(admin.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateTokenAdmins(admin.ID)
}