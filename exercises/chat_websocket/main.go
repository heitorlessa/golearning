package main

import (
	"flag"
	"log"
	"net/http"
	// disabled temporarily as it's not in use
	//"os"
	"path/filepath"
	"sync"
	"text/template"
	// disabled temporarily as it's not in use
	//"trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// Handles HTTP Request (ServeHTTP method)
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	t.once.Do(func() {

		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, r)

}

func main() {

	var addr = flag.String("addr", ":8080", "Address of application.")
	flag.Parse()

	r := newRoom()
	// r.tracer = trace.New(os.Stdout)

	// http.HandleFunc('routeToURL' 'Handler')
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)

	// get the room going (initialize that infinite loop in threads [goroutine])
	go r.run()

	// Log web server startup
	log.Println("Starting web server on...", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
