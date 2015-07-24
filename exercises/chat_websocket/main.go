package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
	"sync"
)

type templateHandler struct {

	once			sync.Once
	filename		string
	templ 			*template.Template
}

// Handles HTTP Request (ServeHTTP method)
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	t.once.Do(func() {
		
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, r)

}

func main() {
	
	r := newRoom()

	// http.HandleFunc('routeToURL' 'Handler')
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle('/room', r)

	// get the room going (initialize that infinite loop in threads [goroutine])
	go r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}