package note

import (
	"log"
	"net/http"
	"strings"

	"github.com/zgr126/x/server/controller/controllerCommon"
	"github.com/zgr126/x/server/model"
)

type Note struct {
}

func New() *Note {
	return &Note{}
}

func (n *Note) CreateNote(w http.ResponseWriter, r *http.Request) {
	note, err := model.InsertNote()
	if err != nil {
		log.Print("err:", err)
	}
	w.Write(controllerCommon.MakeApiJson(note, err))
}

func (n *Note) DeleteNote(w http.ResponseWriter, r *http.Request) {
	noteid := r.FormValue("key")
	err := model.DeleteNote(noteid)
	w.Write(controllerCommon.MakeApiJson(nil, err))
}

func (n *Note) ModifyNote(w http.ResponseWriter, r *http.Request) {
	key := r.PostFormValue("notekey")
	title := r.PostFormValue("title")
	_label := r.PostFormValue("label")
	label := strings.Split(_label, ",")
	content := r.PostFormValue("content")

	note, err := model.ModifyNote(key, title, "", content, label)
	if err != nil {
		log.Print("err:", err)
	}
	w.Write(controllerCommon.MakeApiJson(note, err))
}

func (n *Note) GetNoteLst(w http.ResponseWriter, r *http.Request) {
	notes, err := model.GetAllNote(1, 10, true)
	w.Write(controllerCommon.MakeApiJson(notes, err))
}

func (n *Note) GetNoteContent(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	noteContent, err := model.GetNoteContent(key)
	w.Write(controllerCommon.MakeApiJson(noteContent, err))
}

func (n *Note) SetLabel(w http.ResponseWriter, r *http.Request) {
	noteid := r.FormValue("key")
	err := model.DeleteNote(noteid)
	w.Write(controllerCommon.MakeApiJson(nil, err))
}
