package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type WebServer struct {
	ledController *LedController
}

func NewWebServer(ledController *LedController) *WebServer {
	return &WebServer{ledController: ledController}
}

func (s WebServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func (s WebServer) ActivateLedsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	color := vars["color"]

	_, exist := s.ledController.Leds[color]
	if !exist {
		log.Printf("ERROR: Leds with color '%v' not configured!\n", color)
		w.WriteHeader(http.StatusNotFound)
	} else {
		s.ledController.ActivateLeds(color)
		fmt.Fprintln(w, "OK")
	}
}

func (s WebServer) DeactivateLedsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	color := vars["color"]

	_, exist := s.ledController.Leds[color]
	if !exist {
		log.Printf("ERROR: Leds with color '%v' not configured!\n", color)
		w.WriteHeader(http.StatusNotFound)
	} else {
		s.ledController.DeactivateLeds(color)
		fmt.Fprintln(w, "OK")
	}
}

func (s WebServer) BlinkLedsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	color := vars["color"]

	_, exist := s.ledController.Leds[color]
	if !exist {
		log.Printf("ERROR: Leds with color '%v' not configured!\n", color)
		w.WriteHeader(http.StatusNotFound)
	} else {
		s.ledController.ActivateLeds(color)
		time.Sleep(333 * time.Millisecond)
		s.ledController.DeactivateLeds(color)
		fmt.Fprintln(w, "OK")
	}
}

func (s WebServer) ClearLedsHandler(w http.ResponseWriter, r *http.Request) {
	s.ledController.ClearLeds()
	fmt.Fprintln(w, "OK")
}

func (s WebServer) Run(port int) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", s.IndexHandler)
	router.HandleFunc("/led/{color}/on", s.ActivateLedsHandler)
	router.HandleFunc("/led/{color}/off", s.DeactivateLedsHandler)
	router.HandleFunc("/led/{color}/blink", s.BlinkLedsHandler)
	router.HandleFunc("/leds/clear", s.ClearLedsHandler)

	log.Printf("Starting web server using port: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
