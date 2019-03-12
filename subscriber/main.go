package main

import (
	"fmt"
	"os"
	"html/template"
	"net/http"
	"./mongoDb"
	"./conf"
)

func registry(wr http.ResponseWriter, req *http.Request) {
	fmt.Println("On registry()")
	
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

func main() {
	http.HandleFunc(conf.Conf.Web.Context, registry)
	
	if err := http.ListenAndServe(":" + conf.Conf.Web.Port, nil); err != nil {
		panic(err)
	}
}
