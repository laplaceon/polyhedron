package main

import (
	"github.com/laplaceon/httprouter"
	"fmt"
	"net/http"
)

func DefaultError() string {
	return "error";
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, DefaultError())
}

// takes in a function as a parameter and returns a handlerfunction which can be used by the router
// this way, the original function can still be used when accessed over ZMQ, and the handler can be used when accessed over HTTP
func GenerateHandler(fn (func(*http.Request, httprouter.Params) string)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprint(w, fn(r, ps))
	}
}

func Index(_ *http.Request, _ httprouter.Params) string {
	return "Welcome"
}

func UserFind(r *http.Request, ps httprouter.Params) string {
	var id string = ""
	if(r != nil) {
		r.ParseForm()
		id = r.PostFormValue("id")
	} else {
		id = ps.ByName("id")
	}
	
	if(id == "") {
		return "No id given"
	}
	
	return "Id is " + id
}