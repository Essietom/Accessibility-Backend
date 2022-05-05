package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate *validator.Validate

func ParseBodyTest(r *http.Request, x interface{}, w http.ResponseWriter, ) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
		//	ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
			return
		}
	}
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ValidateInputs(dataSet interface{}) (bool, map[string][]string) {
	validate = validator.New()

	err := validate.Struct(dataSet)

	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		//Validation errors occurred
		errors := make(map[string][]string)
		//Use reflector to reverse engineer struct
		reflected := reflect.ValueOf(dataSet)
		for _, err := range err.(validator.ValidationErrors) {

			//fmt.Println("structf", err.StructField())
			//fmt.Println("refl", reflected.Type().Elem())

			// Attempt to find field by name and get json tag name
			field, _ := reflected.Type().Elem().FieldByName(err.StructField())
			var name string

			//If json tag doesn't exist, use lower case of name
			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = append(errors[name], "The "+name+" is required")
				break
			case "email":
				errors[name] = append(errors[name], "The "+name+" should be a valid email")
				break
			case "eqfield":
				errors[name] = append(errors[name], "The "+name+" should be equal to the "+err.Param())
				break
			default:
				errors[name] = append(errors[name], "The "+name+" is invalid")
				break
			}
		}

		return false, errors
	}
	return true, nil
}


func StringToPrimitive(id string) (primitive.ObjectID){
	primId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		 
	}
	return primId
}

func StringToInt64(str string) (int64){
	n, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("query data type is wrong", str)
	}
	return n
}