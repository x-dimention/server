package router

import (
	"github.com/beego/mux"
	"github.com/zgr126/x/server/controller"
)

var r = mux.New()

func router() {
	r.DefaultHandler(controller.NotMatch)

	r.Get("/", controller.Home)
	r.Post("/", controller.Home)
	r.Get("/abc/:id", controller.Home)

	// note
	r.Post("/api/note/modifyNote", controller.Note.ModifyNote)
	r.Get("/api/note/content", controller.Note.GetNoteContent)
	r.Get("/api/note/getNoteLst", controller.Note.GetNoteLst)
	r.Post("/api/note/deleteNote", controller.Note.DeleteNote)
	r.Get("/api/note/createNote", controller.Note.CreateNote)
	//notelabel
	r.Post("/note/api/setLabel", controller.Note.SetLabel)
}

func GetRouter() *mux.Mux {
	router()
	return r
}
