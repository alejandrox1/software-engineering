package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NotesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "notes"
)

// Establish a connection to database
func (n *NotesDAO) Connect() {
	session, err := mgo.Dial(n.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(n.Database)
}

// Find list of notes.
func (n *NotesDAO) FindAll() ([]Note, error) {
	var notes []Note
	err := db.C(COLLECTION).Find(bson.M{}).All(&notes)
	return notes, err
}

// Find a movie by its id
func (n *NotesDAO) FindById(id string) (Note, error) {
	var note Note
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&note)
	return note, err
}

// Insert a note into database
func (n *NotesDAO) Insert(note Note) error {
	err := db.C(COLLECTION).Insert(&note)
	return err
}

// Update an existing note.
func (n *NotesDAO) Update(note Note) error {
	err := db.C(COLLECTION).UpdateId(note.ID, &note)
	return err
}

// Delete an existing movie
func (n *NotesDAO) Delete(note Note) error {
	err := db.C(COLLECTION).Remove(bson.M{"_id": note.ID})
	return err
}
