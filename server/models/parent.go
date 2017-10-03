package models
import "gopkg.in/mgo.v2/bson"

type Parent struct {
    ID  bson.ObjectId `bson:"_id"`
    Name      string `json:"name"`
    Lastname  string `json:"lastName"`
}

type Parents []Parent
