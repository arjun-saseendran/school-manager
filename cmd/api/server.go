package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/arjun-saseendran/school-manager/internal/api/middlewares"
)

type User struct {
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Place string `json:"place"`
}

func main() {
	port := ":3000"
	cert := "cert.pem"
	key := "key.pem"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("home route"))
		fmt.Println("hello route")
	})

	mux.HandleFunc("/teachers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			fmt.Println("The request url path is: ", r.URL.Path)
			path := strings.TrimPrefix(r.URL.Path, "/teachers/")
			userId := strings.TrimSuffix(path, "/")
			fmt.Println("The user id is: ", userId)
			w.Write([]byte("it is teachers get method."))
			return
		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "failed to  parse form.", http.StatusBadRequest)
				return
			}
			fmt.Println("form: ", r.Form)
			body, err := io.ReadAll(r.Body)
			responseMap := make(map[string]interface{})
			for key, value := range r.Form {
				responseMap[key] = value[0]
			}
			err = json.Unmarshal(body, &responseMap)
			if err != nil {
				http.Error(w, "failed to read the body.", http.StatusBadRequest)
				return
			}
			fmt.Println("Processed response data: ", responseMap)
			w.Write([]byte("it is teachers post method."))

			defer r.Body.Close()
			fmt.Println("The data is: ", body)
			fmt.Println("The string data is: ", string(body))
			var user User
			err = json.Unmarshal(body, &user)
			if err != nil {
				return
			}
			fmt.Println("The user data is: ", user)
			fmt.Println("The user name is: ", user.Name)

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

	mux.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			queryParams := r.URL.Query()
			name := queryParams.Get("name")

			fmt.Println("name is: ", name)
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

	mux.HandleFunc("/exces", func(w http.ResponseWriter, r *http.Request) {
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

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr: port,
		// Handler: middlewares.Cors(mux),

		Handler:   middlewares.Compression(middlewares.ResponseTimeMiddleware(middlewares.SecurityHeaders(middlewares.Cors(mux)))),
		TLSConfig: tlsConfig,
	}

	fmt.Println("server running or port: ", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("failed to start the server.")
	}
}
