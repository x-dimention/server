package hamsterhttp

import (
	"log"
	"net/http"
	"time"
)

type configStruct struct {
	HTTP struct {
		Port string
	}
}

type HttpServe struct {
	Serve  *http.Server
	port   string
	handle http.Handler
}

func New() HttpServe {
	return HttpServe{}
}

func (h *HttpServe) SetHandler(handle http.Handler) {
	h.handle = handle

}

func (h *HttpServe) SetPort(port string) {
	h.port = port
}

func (h *HttpServe) Run() {
	s := &http.Server{
		Addr:           ":" + h.port,
		Handler:        h.handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Print("////////////  OK, Let's Fight!  ////////////")
	log.Fatal(s.ListenAndServe())
}
