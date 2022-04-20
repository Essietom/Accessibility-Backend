package controllers

import (
	"Accessibility-Backend/dto"
	"Accessibility-Backend/entity"
	"Accessibility-Backend/model"
	"Accessibility-Backend/utilities"
	"net/http"

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

	sortByField := r.URL.Query().Get("sortby")
	searchField := r.URL.Query().Get("searchField")
	pageNum := r.URL.Query().Get("pageNum")
	pageSize := r.URL.Query().Get("pageSize")
	orderBy := r.URL.Query().Get("orderBy")
	
	webpageDetails, err := model.GetWebpageByField(searchField, sortByField, orderBy, utilities.StringToInt64(pageSize), utilities.StringToInt64(pageNum))
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
