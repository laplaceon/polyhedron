package main

import (
	"github.com/laplaceon/httprouter"
	"net/http"
)

type Route struct {
	Pattern string
	Method string
	Origin func(*http.Request, httprouter.Params) string
}

type Routes []Route

var routes = Routes{
	Route {
		Pattern: "/",
		Method: "GET",
		Origin: Index,
	},
	Route {
		Pattern: "/user/check",
		Method: "POST",
		Origin: UserFind,
	},
}