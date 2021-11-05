package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cepot-blip/fullstack/api/models"
	"github.com/cepot-blip/fullstack/api/responses"
	"github.com/cepot-blip/fullstack/api/utils/formaterror"
	"github.com/gorilla/mux"
)

//		CREATE BANK
func (server *Server) CreateBank(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bank := models.Bank{}
	err = json.Unmarshal(body, &bank)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bank.Prepare()
	err = bank.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	bankCreated, err := bank.SaveBank(server.DB)

	if err != nil {
		forrmatedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, forrmatedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, bankCreated.ID))
	responses.JSON(w, http.StatusOK, bankCreated)
}

//		READ ALL BANK
func (server *Server) ReadAllBank(w http.ResponseWriter, r *http.Request) {
	bank := models.Bank{}

	banks, err := bank.FindAllBank(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, banks)
}

//		FIND BANK BY ID
func (server *Server) FindBankByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	bank := models.Bank{}
	bankGet, err := bank.FindBankByID(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, bankGet)
}

//		UPDATE BANK
func (server *Server) UpdateBank(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	//		PERIKSA APA ID BANK ITU ADA
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//		PERIKSA APAKAH BANK NYA ADA
	bank := models.Bank{}
	err = server.DB.Debug().Model(models.Bank{}).Where("id = ?", pid).Take(&bank).Error

	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("bank tidak ditemukan"))
		return
	}

	//		BACA DATA YANG DI UPDATE
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//		MUALI MEMPROSES DATA PERMINTAAN
	bankUpdate := models.Bank{}
	err = json.Unmarshal(body, &bankUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

}
