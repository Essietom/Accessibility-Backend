package unitTest

import (
	"Accessibility-Backend/controllers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)


//case1: check for empty name is criteria
func TestCriteriaWithEmptyName(t *testing.T) {

	var jsonString = []byte(`{"name":"", "note":""}`)
	req, _ := http.NewRequest("POST", "criteria", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	//emptyName := controllers.CreateCriteria(,req)
}

//case 2: check for duplicate criteria entry - check by name

func TestGetCriteria(t *testing.T) {
	req, err := http.NewRequest("GET", "/criteria", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetCriteria)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"9","name":"Krish","note":"Bhanushali"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCriteriaById(t *testing.T){
	req, err := http.NewRequest("GET", "/criteria", nil)
	if err!=nil{
		t.Fatal(err)
	}

	q:=req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetCriteriaById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"9","name":"Krish","note":"Bhanushali"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCriteriaByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/criteria", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetCriteriaById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
