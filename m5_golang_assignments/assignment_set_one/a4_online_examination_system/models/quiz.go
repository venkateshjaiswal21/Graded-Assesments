package models

type Question struct {
	Text          string
	Options       []string
	CorrectAnswer int
}
