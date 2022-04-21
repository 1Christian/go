package use_cases

import (
	"crud-app/client"
	"crud-app/entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var person entities.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	client.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
