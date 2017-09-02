package controller

// import (
// 	"golang.org/x/net/context"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func icoController(c context.Context, w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open(Config.Client.Path + "/favicon.ico")
// 	log.Print(Config.Client.Path + "/favicon.ico")
// 	if err != nil {
// 		log.Print("favicon.ico OPEN ERROR!")
// 	}
// 	byt, err := ioutil.ReadAll(f)
// 	if err != nil {
// 		log.Print("favicon.ico READ ERROR!")
// 	}
// 	w.Write(byt)
// }
