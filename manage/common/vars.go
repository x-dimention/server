package common

import (
	"net/http"
)

var CommonVars CommonVarsType
var CV = CommonVars

type CommonVarsType struct {
	HttpPort string
	Router   http.Handler
}

func (v *CommonVarsType) new() {

}
