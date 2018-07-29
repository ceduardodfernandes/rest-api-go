package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
    . "./models"
    . "./dao"
)

var dao = WidgetsDAO{}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.HandleFunc("/widgets", GetWidgets).Methods("GET")
    router.HandleFunc("/widgets/{id}", GetWidget).Methods("GET")
    router.HandleFunc("/widget", CreateWidget).Methods("POST")
    router.HandleFunc("/widget/{id}", PostWidget).Methods("PUT")
    router.HandleFunc("/widget/{id}", DeleteWidget).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dao.FindAllUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := dao.FindUserById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

func GetWidgets(w http.ResponseWriter, r *http.Request) {
	widgets, err := dao.FindAllWidgets()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, widgets)
}

func GetWidget(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	widget, err := dao.FindWidgetById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Widget ID")
		return
	}
	respondWithJson(w, http.StatusOK, widget)
}

func CreateWidget(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	widget.ID = bson.NewObjectId()
	if err := dao.InsertWidget(widget); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, widget)
}

func PostWidget(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateWidget(widget); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteWidget(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var widget Widget
	if err := json.NewDecoder(r.Body).Decode(&widget); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.RemoveWidget(widget); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
