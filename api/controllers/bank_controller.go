package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cepot-blip/fullstack/api/responses"
	"github.com/cepot-blip/fullstack/models"
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
	

}