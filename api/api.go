package api

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type WebhookBodyReq struct {
	Object string     `json:"object"`
	Entry  []EntryReq `json:"entry"`
}

type EntryReq struct {
	Messaging []MessagingReq `json:"messaging"`
}

type MessagingReq struct {
	Message string `json:"message"`
}

//Start init http apis
func Start() {
	log.Info("initial api setup")

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/webhook", webhookPostHandle).Methods("POST")
	http.Handle("/", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func webhookPostHandle(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	var req WebhookBodyReq
	if err := d.Decode(&req); err != nil {
		log.Error("error when decode request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Infof("body: %s", req)
}
