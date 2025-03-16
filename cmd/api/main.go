package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/timooo-thy/go_tutorial/cmd/internal/handlers"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Server is running on port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
}