package apiserver

import (
	"api-go/apiserver/get_pictures"
	"api-go/apiserver/login"
	"api-go/apiserver/upload_picture"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

// creating config
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	server.ConfigRouter()

	fmt.Println("Starting API-Server " + server.config.BindAddr)

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func Decode_Request(req *http.Request, data interface{}) {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	dec.Decode(data)
}

// configuring router: applying hadlers to routes
func (server *APIServer) ConfigRouter() {
	server.router.HandleFunc("/login", server.HandleLogin())
	server.router.HandleFunc("/upload-picture", server.HandleUploadPicture())
	server.router.HandleFunc("/get-pictures", server.HandleGetPictures())
}

// handlers for routes
func (server *APIServer) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				var data login.User
				Decode_Request(r, &data)

				response, successful := login.Login(data)
				response_jsonify, _ := json.Marshal(response)

				if !successful {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.Write(response_jsonify)
			}
		default:
			{
				w.Write([]byte("Not supporting this method"))
			}
		}
	}
}

func (server *APIServer) HandleUploadPicture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				tokenString := r.Header.Get("Authorization")
				tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

				var data upload_picture.UploadPicture
				Decode_Request(r, &data)

				response, successful := upload_picture.Upload_Picture(data, tokenString)
				response_jsonify, _ := json.Marshal(response)

				if !successful {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.Write(response_jsonify)
			}
		default:
			{
				w.Write([]byte("Not supporting this method"))
			}
		}
	}
}

func (server *APIServer) HandleGetPictures() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				tokenString := r.Header.Get("Authorization")
				tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

				response, successful := get_pictures.Get_Pictures(tokenString)
				response_jsonify, _ := json.Marshal(response)

				if !successful {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.Write(response_jsonify)
			}
		default:
			{
				w.Write([]byte("Not supporting this method"))
			}
		}
	}
}
