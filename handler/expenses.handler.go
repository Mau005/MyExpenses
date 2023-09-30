package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Mau005/MyExpenses/configuration"
	"github.com/Mau005/MyExpenses/controller"
	"github.com/Mau005/MyExpenses/models"
	"github.com/gorilla/mux"
)

type ExpensesHandler struct{}

func (ex *ExpensesHandler) CreateExpensesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entra a la consulta?")
	var exp controller.ExpensesController

	name := r.FormValue("name")
	nameCategory := r.FormValue("namecategory")
	description := r.FormValue("description")

	var cat controller.CategoryController

	category, err := cat.GetCategoryName(nameCategory)
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

	expense := models.Expenses{Name: name,
		Category:    category,
		Description: description}

	expense, err = exp.CreateExpenses(expense)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(expense)
}

func (ex *ExpensesHandler) GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ExpensesController

	content := mux.Vars(r)
	idExpenses, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	expenses, err := exp.GetExpenses(uint(idExpenses))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(expenses)
}

func (ex *ExpensesHandler) GetAllExpensesHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ExpensesController

	expenses, err := exp.GetAllExpenses()

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(expenses)
}

func (ex *ExpensesHandler) PatchExpensesHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	var exp controller.ExpensesController

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	expenses, err := exp.PuthExpense(data)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}
	json.NewEncoder(w).Encode(expenses)

}

func (ex *ExpensesHandler) DelExpensesHandler(w http.ResponseWriter, r *http.Request) {
	var exp controller.ExpensesController

	content := mux.Vars(r)
	expenses, err := strconv.ParseUint(content["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
			Status:    http.StatusNotAcceptable,
			Message:   err.Error(),
			TimeStamp: time.Now(),
		})
		return
	}

	err = exp.DeleteExpense(uint(expenses))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(models.Exception{
			Error:     configuration.ERROR_SERVICE_EXPENSES,
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
		Message: configuration.DELETE_EXPENSES,
		Status:  http.StatusAccepted,
	})
}
