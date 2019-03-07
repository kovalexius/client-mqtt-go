package main

import (
	//"fmt"
	"net/http"
	_ "./mongo-db"
)

func registry(wr http.ResponseWriter, req *http.Request) {
	
	path := "static/img" + req.URL.Path
	if file, err := os.Stat(path); err == nil && !file.IsDir() {
		svg := template.Must(template.ParseFiles(path))
	
		data := struct {
			Color string
			Use   string
		}{}

		color := req.URL.Query().Get("c")
		if n := len(color); n != 0 && n < 7 {
			data.Color = color
		}
		use := req.URL.Query().Get("u")
		if len(use) != 0 {
			data.Use = use
		}

		if err := svg.Execute(wr, data); err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.NotFound(wr, req)
}

func main() {
	http.Handle("/registry", registry)
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
