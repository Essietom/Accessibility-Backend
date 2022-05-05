package controllers

import (
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

var NewCriteria entity.Criteria

func CreateCriteria(w http.ResponseWriter, r *http.Request) {
	// utilities.EnableCors(&w)

	vi := &entity.Criteria{}
	utilities.ParseBody(r, vi)
	v, _ := model.CreateCriteria(vi)
	utilities.SuccessRespond(v, w, r)
}

func GetCriteria(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	criteria, err := model.GetAllCriteria()
	if err != nil {

	}
	utilities.SuccessRespond(criteria, w, r)
}

func GetCriteriaById(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]
	criteriaDetails, err := model.GetCriteriaById(criteriaId)
	if err != nil {

	}
	utilities.SuccessRespond(criteriaDetails, w, r)
}

func UpdateCriteria(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	var updateCriteria = &entity.Criteria{}
	utilities.ParseBody(r, updateCriteria)
	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]

	criteriaDetails, err := model.GetCriteriaById(criteriaId)
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
	model.UpdateCriteria(criteriaDetails, criteriaId)
	utilities.SuccessRespond(criteriaDetails, w, r)

}

func DeleteCriteria(w http.ResponseWriter, r *http.Request) {
	//utilities.EnableCors(&w)

	vars := mux.Vars(r)
	criteriaId := vars["criteriaId"]
	primitiveId, _ := primitive.ObjectIDFromHex(criteriaId)
	vin := model.DeleteCriteria(primitiveId)
	utilities.SuccessRespond(vin, w, r)

}
