package main

import (
    "encoding/json"
    "fmt"
    "github.com/server/models"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

type Handler struct{
  Repository Repository
}

func (c *Handler)Index(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintln(w, "Welcome!")
    parents := c.Repository.GetParents() // list of all albums
    log.Println(parents)
    data, _ := json.Marshal(parents)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

func ParentIndex(w http.ResponseWriter, r *http.Request) {
    parent := models.Parent{
      Name: "jos",
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(parent); err != nil {
        panic(err)
    }

    //json.NewEncoder(w).Encode(u)
}

func ParentShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    parentId := vars["parentId"]
    fmt.Fprintln(w, "Parent show:", parentId)
}
