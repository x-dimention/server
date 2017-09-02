package common

import (
	"net/http"

	"github.com/beego/mux"
)

var commonHttpMethod commonHttpMethodType
var CHM = commonHttpMethod

type commonHttpMethodType struct {
}

func Params(r *http.Request) map[string]string {
	return mux.Params(r)
}

// Param return the router param based on the key
func Param(r *http.Request, key string) string {
	return mux.Param(r, key)
}
