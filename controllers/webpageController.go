package controllers

import (
	"Accessibility-Backend/entity"
	"Accessibility-Backend/models"
	"Accessibility-Backend/utilities"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var Webpage entity.Webpage

func SaveWebpageScans(w http.ResponseWriter, r *http.Request) {
	vi := &entity.Webpage{}
	utilities.ParseBody(r, vi)
	if ok, errors := utilities.ValidateInputs(vi); !ok {
		utilities.ValidationResponse(errors, w)
		return
	}
	v, err := models.SaveWebpageScan(vi)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(v, w)

}

func GetWebpageScan(w http.ResponseWriter, r *http.Request) {
	webpageScan, err := models.GetAllWebpages()
	if err != nil {

	}
	res, _ := json.Marshal(webpageScan)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetWebpageByField(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageField := vars["webpageField"]
	webpageDetails, err := models.GetWebpageByField(webpageField)
	if err != nil {

	}
	res, _ := json.Marshal(webpageDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetWebpageScanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	webpageDetails, err := models.GetWebpageById(utilities.StringToPrimitive(webpageId))
	if err != nil {

	}
	res, _ := json.Marshal(webpageDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateWebpageScan(w http.ResponseWriter, r *http.Request) {
	var updateWebpage = &entity.Webpage{}
	utilities.ParseBody(r, updateWebpage)
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]

	webpageDetails, err := models.GetWebpageById(utilities.StringToPrimitive(webpageId))
	if err != nil {
		println("tomiiiii", err)
		return

	}

	if updateWebpage.Name != "" {
		webpageDetails.Name = updateWebpage.Name
	}
	if updateWebpage.Note != "" {
		webpageDetails.Note = updateWebpage.Note
	}
	models.UpdateWebpage(webpageDetails, webpageId)
	res, _ := json.Marshal(webpageDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteWebpageScan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	primitiveId, _ := primitive.ObjectIDFromHex(webpageId)
	vin := models.DeleteWebpage(primitiveId)
	res, _ := json.Marshal(vin)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
