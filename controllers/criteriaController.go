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

var NewCriteria entity.Criteria

func CreateCriteria(w http.ResponseWriter, r *http.Request) {
	vi := &entity.Criteria{}
	utilities.ParseBody(r, vi)
	v, _ := models.CreateCriteria(vi)
	res, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCriteria(w http.ResponseWriter, r *http.Request) {
	criteria, err := models.GetAllCriteria()
	if err != nil {

	}
	res, _ := json.Marshal(criteria)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCriteriaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]
	criteriaDetails, err := models.GetCriteriaById(criteriaId)
	if err != nil {

	}
	res, _ := json.Marshal(criteriaDetails)
	w.Header().Set("Content-Type", "pkglicatio/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCriteria(w http.ResponseWriter, r *http.Request) {
	var updateCriteria = &entity.Criteria{}
	utilities.ParseBody(r, updateCriteria)
	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]

	criteriaDetails, err := models.GetCriteriaById(criteriaId)
	if err != nil {
		println("tomiiiii", err)
		return

	}

	if updateCriteria.Name != "" {
		criteriaDetails.Name = updateCriteria.Name
	}
	if updateCriteria.Note != "" {
		criteriaDetails.Note = updateCriteria.Note
	}
	models.UpdateCriteria(criteriaDetails, criteriaId)
	res, _ := json.Marshal(criteriaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteCriteria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]
	primitiveId, _ := primitive.ObjectIDFromHex(criteriaId)
	vin := models.DeleteCriteria(primitiveId)
	res, _ := json.Marshal(vin)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
