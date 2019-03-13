package main

import (
	"fmt"
	"os"
	"html/template"
	"net/http"
	"./mongoDb"
	"./conf"
)

func get(wr http.ResponseWriter, req *http.Request) {
	path := conf.Conf.Web.HtmlStatic
	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		page := template.Must(template.ParseFiles(path))
		
		records := mongoDb.GetTelemetry()

		if err := page.Execute(wr, records); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.NotFound(wr, req)
}

func post(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("POST Method")
}

func put(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("PUT Method")
}

func delete(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("DELETE Method")
}

func registry(wr http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":
			get(wr, req)
		case "POST":
			post(wr, req)
		case "PUT":
			put(wr, req)
		case "DELETE":
			delete(wr, req)
		default:
			http.Error(wr, "method " + req.Method + " not supported", 405)
	}
}

func main() {
	http.HandleFunc(conf.Conf.Web.Context, registry)
	
	if err := http.ListenAndServe(":" + conf.Conf.Web.Port, nil); err != nil {
		panic(err)
	}
}
