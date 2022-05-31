package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Webpage entity.Webpage


func SaveWebpageScans(w http.ResponseWriter, r *http.Request) {
	webpageRequest := &dto.WebpageRequestBody{}

	 err := utilities.ParseBodyTest(r, webpageRequest, w)
	 if err != nil {
		utilities.ErrorResponse(http.StatusBadRequest, err.Error(), w,r)
		return
	}
	if ok, errors := utilities.ValidateInputs(webpageRequest); !ok {
		utilities.ValidationResponse(errors, w,r)
		return
	}
	v, err := model.SaveWebpageScan(webpageRequest)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w,r)
		return
	}
	utilities.SuccessRespond(v, w,r)

}

func GetWebpageScan(w http.ResponseWriter, r *http.Request) {

	webpageScan, err := model.GetAllWebpages()
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w,r)
		return
	}
	utilities.SuccessRespond(webpageScan, w,r)
}

func GetWebpageByField(w http.ResponseWriter, r *http.Request) {

	sortByField := r.URL.Query().Get("sortBy")
	searchField := r.URL.Query().Get("searchField")
	pageNum := r.URL.Query().Get("pageNum")
	pageSize := r.URL.Query().Get("pageSize")
	orderBy := r.URL.Query().Get("orderBy")

	var err error

	orderByInt := 1
	if orderBy != "" && orderBy == "asc" {
		orderByInt = 1
	} else {
		orderByInt = -1
	}

	sortQuery := "scantime"
	if sortByField != ""{
	sortQuery, err = validateAndReturnSortQuery(sortByField)
	if err != nil {
		utilities.ErrorResponse(400, err.Error(), w, r)
		return
	}
	}
	

	limit := 10
	if pageSize != "" {
		limit, err = strconv.Atoi(pageSize)
		if limit < 1 {
			utilities.ErrorResponse(400, "invalid value for page size", w, r)
			return
		}
		if err != nil{
			utilities.ErrorResponse(500, err.Error(), w, r)
			return
		}
	}

	pgnum := 1
	if pageNum != "" {
		pgnum, err = strconv.Atoi(pageNum)
		if limit < -1 {
			utilities.ErrorResponse(400, "invalid value for page size", w, r)
			return
		}
		if err != nil{
			utilities.ErrorResponse(400, err.Error(), w, r)
			return		
		}
	}
	log.Print("fields", searchField, sortQuery, orderByInt)
	webpageDetails, err := model.GetWebpageByField(searchField, sortQuery, orderByInt, int64(limit), int64(pgnum))
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(webpageDetails, w, r)
}

func GetWebpageScanById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	wpid, err := primitive.ObjectIDFromHex(webpageId)
	if err != nil {
		utilities.ErrorResponse(400, "the webpageid passed is invalid", w, r)
		return
	}
	webpageDetails, err := model.GetWebpageById(wpid)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond(webpageDetails, w, r)
}

// func UpdateWebpageScan(w http.ResponseWriter, r *http.Request) {
// 	var updateWebpage = &entity.Webpage{}
// 	utilities.ParseBody(r, updateWebpage)
// 	vars := mux.Vars(r)
// 	webpageId := vars["webpageId"]

// 	webpageDetails, err := model.GetWebpageById(webpageId)
// 	if err != nil {
// 		println("tomiiiii", err)
// 		return

// 	}

// 	if updateWebpage.Name != "" {
// 		webpageDetails.Name = updateWebpage.Name
// 	}
// 	if updateWebpage.Note != "" {
// 		webpageDetails.Note = updateWebpage.Note
// 	}
// 	model.UpdateWebpage(webpageDetails, webpageId)
// 	res, _ := json.Marshal(webpageDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

func DeleteWebpageScan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	objectId, err := primitive.ObjectIDFromHex(webpageId)
	if err != nil {
		utilities.ErrorResponse(400, "the webpageid passed is invalid", w, r)
		return
	}
	err = model.DeleteWebpage(objectId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w, r)
		return
	}
	utilities.SuccessRespond("sucessfully deleted", w, r)

}


func stringInSlice(strSlice []string, s string) bool {
	for _, v := range strSlice {
		if v == s {
			return true
		}
	}

	return false
}

func getWebpageFields() []string {
	var field []string

	v := reflect.ValueOf(entity.Webpage{})
	for i := 0; i < v.Type().NumField(); i++ {
		field = append(field, v.Type().Field(i).Tag.Get("json"))
	}

	return field
}

func validateAndReturnSortQuery(sortBy string) (string, error) {
	
	if !stringInSlice([]string{"name", "scantime", "website", "url"}, sortBy) {
		return "", errors.New("unknown field in sortBy query parameter")
	}

	return sortBy, nil

}

