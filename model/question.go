package models

//Question asked by users
type Question struct {
	UserEntity
	Text     string
	AnswerID UniqueID
}