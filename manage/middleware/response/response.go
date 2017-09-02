package response

import (
	"net/http"
)

type Response struct {
}

// Middleware is a struct that has a ServeHTTP method
func NewMiddleware() *Response {
	return &Response{}
}

// The middleware handler
func (l *Response) ServeHTTP(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// Log moose status
	// log.Println("req:" + req.Body)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call the next middleware handler
	next(w, req)
}
