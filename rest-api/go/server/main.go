package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = NotesDAO{}

// GET list of notes.
func AllNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, notes)
}

// GET a note by its ID.
func FindNoteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	note, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, note)
}

// POST a new note.
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	note.ID = bson.NewObjectId()
	if err := dao.Insert(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, note)
}

// PUT update an existing note.
func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	params := mux.Vars(r)
	note.ID = bson.ObjectIdHex(params["id"])
	if err := dao.Update(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing note.
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	params := mux.Vars(r)
	note.ID = bson.ObjectIdHex(params["id"])
	if err := dao.Delete(note); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", AllNotesHandler).Methods("GET")
	r.HandleFunc("/notes", CreateNoteHandler).Methods("POST")
	r.HandleFunc("/notes/{id}", UpdateNoteHandler).Methods("PUT")
	r.HandleFunc("/notes/{id}", FindNoteHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
