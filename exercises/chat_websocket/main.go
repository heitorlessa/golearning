package main

import (
	"log"
	"net/http"
)

func main() {
	
	// http.HandleFunc('routeToURL' 'Handler')
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		w.Write([]byte(`
			<html>
			<head>
				<title>Chat - Sample</title>
			</head>
			<body>
				<div>
					Let's chat
				</div>
			</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}