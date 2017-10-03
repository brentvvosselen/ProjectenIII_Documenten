package main

import (
    "net/http"
)

var controller = &Handler{Repository: Repository{}}

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        controller.Index,
    },
    Route{
        "ParentIndex",
        "GET",
        "/parent",
        ParentIndex,
    },
    Route{
        "ParentShow",
        "GET",
        "/parent/{parentId}",
        ParentShow,
    },
}
