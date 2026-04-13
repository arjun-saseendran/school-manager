package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("home route"))
		fmt.Println("hello route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("it is teachers get method."))
			return
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "failed to  parse form.", http.StatusBadRequest)
			}
			fmt.Println("form: ", r.Form)
			w.Write([]byte("it is teachers post method."))
			return
		case http.MethodPut:
			w.Write([]byte("it is teachers put method"))
			return
		case http.MethodDelete:
			w.Write([]byte("it is teachers delete method"))
			return
		default:
			w.Write([]byte("invalid request"))
		}

	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("it is students get method."))
			return
		case http.MethodPost:
			w.Write([]byte("it is students post method."))
			return
		case http.MethodPut:
			w.Write([]byte("it is students put method"))
			return
		case http.MethodDelete:
			w.Write([]byte("it is students delete method"))
			return
		default:
			w.Write([]byte("invalid request"))
		}

	})

	http.HandleFunc("/exces", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("it is exces get method."))
			return
		case http.MethodPost:
			w.Write([]byte("it is exces post method."))
			return
		case http.MethodPut:
			w.Write([]byte("it is exces put method"))
			return
		case http.MethodDelete:
			w.Write([]byte("it is exces delete method"))
			return
		default:
			w.Write([]byte("invalid request"))
		}

	})

	fmt.Println("server running or port: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("failed to start the server.")
	}
}
