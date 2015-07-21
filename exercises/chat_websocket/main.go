package main

import (
	"log"
	"net/http"
	"text/template"
	"path/filepath"
)

type templateHandler struct {

	once			sync.Once
	filename		string
	templ 			*template.Template
}

func (t *templatehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	t.once.Do(func() {
		
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)

}

func main() {
	
	// http.HandleFunc('routeToURL' 'Handler')
	http.HandleFunc("/", &templateHandler{filename: "chat.html"})
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}