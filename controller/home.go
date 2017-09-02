package controller

import (
	"net/http"

	"github.com/zgr126/x/server/manage/common"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("xblog" + common.CV.HttpPort))
}

// func home(c context.Context, w http.ResponseWriter, r *http.Request) {
// 	// o := struct {
// 	// 	A string   `json:"a"`
// 	// 	B []string `json:"b"`
// 	// }{
// 	// 	A: "6",
// 	// 	B: []string{"g", "s"},
// 	// }
// 	// result, err := model.Orm().Insert(o).Collection("page")
// 	// log.Print(result, "//", err)
// 	// result, err = model.Orm().Query(`
// 	// 	FOR u IN page
// 	//    	RETURN u
// 	// 	`).Collection("page")
// 	// log.Print(result, "//", err)
// 	test.RangePrint(r.Header)
// 	w.Write([]byte("hello"))
// }
