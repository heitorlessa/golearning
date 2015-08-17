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
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
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

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)

}

func main() {

	var addr = flag.String("addr", ":8080", "Address of application.")
	flag.Parse()

	// set up gomniauth
	// callback URLs that will receive auth token comes as 3rd argument for each provider
	gomniauth.SetSecurityKey("some key here")
	gomniauth.WithProviders(
		facebook.New("key", "secret",
			"http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret",
			"http://localhost:8080/auth/callback/github"),
		google.New("key", "secret",
			"http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	// r.tracer = trace.New(os.Stdout)

	// http.Handle('routeToURL' 'Handler')
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))

	// No need for MustAuth wrapper otherwise it goes to an infinite redirection loop
	http.Handle("/login", &templateHandler{filename: "login.html"})

	// Since we don't need to maintain any state (object) we can use HandleFunc and pass a function to it
	http.HandleFunc("/auth/", loginHandler)

	http.Handle("/room", r)

	// get the room going (initialize that infinite loop in threads [goroutine])
	go r.run()

	// Log web server startup
	log.Println("Starting web server on...", *addr)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
