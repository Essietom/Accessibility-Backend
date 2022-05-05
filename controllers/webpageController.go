package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var Webpage entity.Webpage

func SaveWebpageScans(w http.ResponseWriter, r *http.Request) {
	webpageRequest := &dto.WebpageRequestBody{}
	utilities.ParseBody(r, webpageRequest)
	if ok, errors := utilities.ValidateInputs(webpageRequest); !ok {
		utilities.ValidationResponse(errors, w)
		return
	}
	v, err := model.SaveWebpageScan(webpageRequest)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(v, w)

}

func GetWebpageScan(w http.ResponseWriter, r *http.Request) {
	webpageScan, err := model.GetAllWebpages()
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(webpageScan, w)
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

	sortQuery := ""
	if sortByField != ""{
	sortQuery, err = validateAndReturnSortQuery(sortByField)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	}
	

	limit := -1
	if pageSize != "" {
		limit, err = strconv.Atoi(pageSize)
		if limit < -1 {
			utilities.ErrorResponse(500, "invalid value for page size", w)
			return
		}
		if err != nil{
			utilities.ErrorResponse(500, err.Error(), w)
			return
		}
	}

	pgnum := 1
	if pageNum != "" {
		pgnum, err = strconv.Atoi(pageNum)
		if limit < -1 {
			utilities.ErrorResponse(500, "invalid value for page size", w)
			return
		}
		if err != nil{
			utilities.ErrorResponse(500, err.Error(), w)
			return		
		}
	}
	
	webpageDetails, err := model.GetWebpageByField(searchField, sortQuery, orderByInt, int64(limit), int64(pgnum))
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(webpageDetails, w)
}

func GetWebpageScanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	webpageId := vars["webpageId"]
	webpageDetails, err := model.GetWebpageById(webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond(webpageDetails, w)
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
	err := model.DeleteWebpage(webpageId)
	if err != nil {
		utilities.ErrorResponse(500, err.Error(), w)
		return
	}
	utilities.SuccessRespond("sucessfully deleted", w)

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
	
	if !stringInSlice(getWebpageFields(), sortBy) {
		return "", errors.New("unknown field in sortBy query parameter")
	}

	return sortBy, nil

}

func validateAndReturnFilterMap(filter string) (map[string]string, error) {
	splits := strings.Split(filter, ".")
	if len(splits) != 2 {
		return nil, errors.New("malformed sortBy query parameter, should be field.orderdirection")
	}

	field, value := splits[0], splits[1]

	if !stringInSlice(getWebpageFields(), field) {
		return nil, errors.New("unknown field in filter query parameter")
	}

	return map[string]string{field: value}, nil
}