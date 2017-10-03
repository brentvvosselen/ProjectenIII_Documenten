package main

import (
 "fmt"
 "log"
 "github.com/server/models"
 "gopkg.in/mgo.v2"
 "gopkg.in/mgo.v2/bson"
 
)

//Repository ...
type Repository struct{}
// SERVER the DB server
const SERVER = "localhost:27017"
// DBNAME the name of the DB instance
const DBNAME = "heenenweer"
// DOCNAME the name of the document
const DOCNAME = "parents"

// GetParents returns a list of parents
func (r Repository) GetParents() models.Parents {
  session, err := mgo.Dial(SERVER)
  if err != nil {
    fmt.Println("Failed to establish connection to Mongo server:", err)
  }
  defer session.Close()
  c := session.DB(DBNAME).C(DOCNAME)
  results := models.Parents{}
  if err := c.Find(nil).All(&results); err != nil {
    fmt.Println("Failed to write results:", err)
  }
  return results
}

// AddParents adds a parent in the db
func (r Repository) AddParent(parent models.Parent) bool {
 session, err := mgo.Dial(SERVER)
 defer session.Close()
 parent.ID = bson.NewObjectId()
 session.DB(DBNAME).C(DOCNAME).Insert(parent)
 if err != nil {
   log.Fatal(err)
   return false
 }
 return true
}

// UpdateParent updates a parent(not used for now)
func (r Repository) UpdateParent(parent models.Parent) bool {
 session, err := mgo.Dial(SERVER)
 defer session.Close()
 parent.ID = bson.NewObjectId()
 session.DB(DBNAME).C(DOCNAME).UpdateId(parent.ID, parent)
if err != nil {
  log.Fatal(err)
  return false
 }
 return true
}

// DeleteAlbum deletes an Album (not used for now)
func (r Repository) DeleteParent(id string) string {
 session, err := mgo.Dial(SERVER)
 defer session.Close()
// Verify id is ObjectId, otherwise bail
 if !bson.IsObjectIdHex(id) {
  return "NOT FOUND"
 }
// Grab id
 oid := bson.ObjectIdHex(id)
// Remove user
 if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
  log.Fatal(err)
  return "INTERNAL ERR"
 }
// Write status
 return "OK"
}
