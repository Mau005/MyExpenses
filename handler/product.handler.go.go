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

type ProductHandler struct{}

func (ex *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ProductController

	name := r.FormValue("name")
	description := r.FormValue("description")
	ammount := r.FormValue("ammount")
	price := r.FormValue("price")
	expenseID := r.FormValue("expensesid")

	var ac controller.ApiController
	product := models.Product{
		Name:        name,
		Description: description,
		Ammount:     ac.ParseUintDefault(ammount, 1),
		Price:       ac.ParseUintDefault(price, 1),
		ExpensesID:  ac.ParseUintDefault(expenseID, 0),
	}

	product, err := exp.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (ex *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ProductController

	content := mux.Vars(r)
	idProduct, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	product, err := exp.GetProduct(uint(idProduct))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (ex *ProductHandler) GetAllProductHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ProductController

	Product, err := exp.GetAllProduct()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(Product)
}

func (ex *ProductHandler) PatchProductHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var exp controller.ProductController

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	product, err := exp.PutchProduct(data)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(product)

}

func (ex *ProductHandler) DelProductHandler(w http.ResponseWriter, r *http.Request) {
	var pc controller.ProductController

	content := mux.Vars(r)
	Product, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	err = pc.DelProduct(uint(Product))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_PRODUCT,
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
		Message: configuration.DELETE_PRODUCT,
		Status:  http.StatusAccepted,
	})
}
