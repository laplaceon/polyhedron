Polyhedron
======

Polyhedron is a(n) (experimental) framework (?) for developing APIs in Go. It's named so because of the more than 1 ways you can access the API.

Polyhedron can be used to implement an architecture in which the web app (or any client for that matter) consumes its own API, the so called apicentric or api-first design.

When accessed over HTTP, Polyhedron serves normally. It can also be accessed through sockets using ZeroMQ. This communication must be done through a client wrapper which can be written in a multitude of languages. You can either write your own for your app or use prewritten clients which can be found here: <https://github.com/laplaceon/polyhedron-clients>

Usage
-----------
Polyhedron has only been tested on Linux (Ubuntu 14.04). It will probably work on Mac OSX and Windows. Although for Windows, some changes need to be made.

The prerequisites for use are
* Golang installed
* ZeroMQ installed

Use ```go build ./api/*.go``` to compile an executable app file. This assumes the project files are in a directory named api.

Windows
-----------
By default, Polyhedron uses UNIX domain sockets for communication. You will need to change this to communicating over TCP sockets when on a Windows machine.
