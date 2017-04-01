package main

import (
	"fmt"
	"log"
	"net/http"
)

// MapServer fulfills the net Handler inferface
type MapServer map[string]int

// ServeHttp fulfills Handler for MapServer
func (m MapServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	k := req.RequestURI
	if _, ok := m[k]; ok {
		m[k]++
	} else {
		m[k] = 1
	}
	fmt.Fprintf(res, "%+v", m)
	return
}

func main() {
	mapserver := MapServer{}
	log.Println("Serving on :5555")
	http.ListenAndServe(":5555", mapserver)
}
