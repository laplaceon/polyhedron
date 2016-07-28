package main

import (
	"net/http"
	"github.com/laplaceon/httprouter"
	zmq "github.com/pebbe/zmq4"
	"fmt"
	"encoding/json"
	"os"
)

var router *httprouter.Router

type ClientMessage struct {
	Method string
	Data string
	Parameters map[string]interface{}
}

func ListenAndServeZMQ() {
	zmqRouter, _ := zmq.NewSocket(zmq.ROUTER)
	defer zmqRouter.Close()
	
	// inprocess, untested
	// zmqRouter.Bind("inproc://air")
	
	// ipc socket, preferred
	zmqRouter.Bind("ipc:///tmp/air.ipc")
	// make socket writable by client
	os.Chmod("/tmp/air.ipc", 0666)
	
	// tcp socket, works but ipc is a much better option when local
	// zmqRouter.Bind("tcp://*:5555")

	for {
		id, _ := zmqRouter.Recv(0)
		zmqRouter.Recv(0)
		message, _ := zmqRouter.Recv(0)
		if(id != "") {
			// three part envelope
			zmqRouter.Send(id, zmq.SNDMORE)
			zmqRouter.Send("", zmq.SNDMORE)
			zmqRouter.Send(RouteApiResponse(message), 0)
		}
	}
}

func RouteApiResponse(request string) string {
	// create clientmessage and convert request into parsable format
	var clientMessage ClientMessage
	
	if(json.Unmarshal([]byte(request), &clientMessage) != nil) {
		fmt.Println("Error reading response");
	}
	
	path, params, _ := router.LookupPath(clientMessage.Method, clientMessage.Data)
	
	// find handler method associated with request and then execute
	// path, params, _ := router.LookupPath(clientMessage.Method, clientMessage.Data)
	
	var f func(*http.Request, httprouter.Params) string = nil
	
	for i := 0; i < len(routes); i++ {
		if(routes[i].Pattern == path) {
			f = routes[i].Origin
		}
	}

	var response string
	
	if(f != nil) {
		// insert parameters into http params
		for k, v := range clientMessage.Parameters {
			params = append(params, httprouter.Param{Key: k, Value: fmt.Sprintf("%v", v)})
		}
		response = f(nil, params)
	} else {
		response = DefaultError()
	}
	
	if(response == "") {
		// set default response
		response = "error"
	}

	return response
}

func main() {
	router = httprouter.New()
	
	for i := 0; i < len(routes); i++ {
		router.Handle(routes[i].Method, routes[i].Pattern, GenerateHandler(routes[i].Origin))
	}

	// error handler
	router.NotFound = http.HandlerFunc(NotFound);

	// start zeromq router in new goroutine
	go ListenAndServeZMQ()

	port := "8080"

	fmt.Println("Started on port " + port)
	http.ListenAndServe(":" + port, router)
}