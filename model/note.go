package model

import (
	"github.com/zgr126/x-tool"
)

type Note struct {
	Label        []string `xorm:"jsonb"`
	Author       string
	Title        string
	Introduction string
	Content      string
	Views        int64
}

type Note_lst struct {
	Model
	Label []string
	Title string
}

type Note_rec struct {
	ModeRec
	Note
}

type NoteLabel struct {
	Model     `xorm:"extends"`
	Labelname string
	Parent    string
}

// type NoteMark struct {
// 	NoteKey  string
// 	LikerId  string
// 	AuthorId string
// }

const NoteColName = "Note"

func InsertNote() (*Note, error) {
	note := &Note{
		Key:          tool.NewUniqueId(),
		Title:        "NEW NOTE",
		Introduction: "",
		Content:      "",
		Label:        []string{"aaaa", "bbbb"},
	}
	_, err := DB.InsertOne(note)
	if err != nil {
		return note, err
	}
	return note, nil
}

func DeleteNote(Key string) error {
	var note Note
	note.Key = Key
	_, err := DB.Delete(&note)
	return err
}

func ModifyNote(Key, title, introduction, content string, labels []string) (*Note, error) {
	note := &Note{
		Title:        title,
		Introduction: introduction,
		Label:        labels,
		Content:      content,
	}
	DB.Where("Key=?", Key).Cols("title", "introduction", "label", "content").Update(note)
	return note, nil
}

func GetAllNote(page, cont int, desc bool) (noteLst []*Note, err error) {
	DB.Find(&noteLst)
	// noteLst = append(noteLst, &note)
	return
}

func GetNoteContent(key string) (string, error) {
	var note Note
	err := DB.Where("key=?", key).Find(&note)
	return note.Content, err
}

func returnNote(*Note) {

}

func CreateNoteLabel() {

}
