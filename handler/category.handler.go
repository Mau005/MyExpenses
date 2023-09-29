package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/controller"
	"github.com/Mau005/MyExpenses/models"
	"github.com/gorilla/mux"
)

type Categoryhandler struct{}

func (ch *Categoryhandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var cc controller.CategoryController

	name := r.FormValue("name")

	category, err := cc.CreateCategory(name)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (ch *Categoryhandler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var cc controller.CategoryController

	content := mux.Vars(r)
	idCategory, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	category, err := cc.GetCategory(uint(idCategory))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (ch *Categoryhandler) GetAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var cc controller.CategoryController

	categorys, err := cc.GetAllCategory()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(categorys)
}

func (ch *Categoryhandler) PatchCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var cc controller.CategoryController

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	category, err := cc.PutchCategory(data)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(category)

}

func (ch *Categoryhandler) DelCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var cc controller.CategoryController

	content := mux.Vars(r)
	idCategory, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	err = cc.DelCategory(uint(idCategory))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_CATEGORY,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(struct {
		Message string
		Status  int
	}{
		Message: configuration.DELETE_CATEGORY,
		Status:  http.StatusAccepted,
	})
}
