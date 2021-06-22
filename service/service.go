package service

import (
	"fmt"
	"github.com/wambugucoder/golang-graphql/graph/model"
	"github.com/wambugucoder/golang-graphql/repository"
)

func CreateQuestion(input model.QuestionInput) (*model.Question, error) {
	fmt.Println("input", input.QuestionText, input.PubDate)
	question := model.Question{}
	question.ID = "1"
	question.QuestionText = input.QuestionText
	question.PubDate = input.PubDate
	return repository.CreateQuestion(question)
}

func CreateChoice(input *model.ChoiceInput) (*model.Choice, error) {
	choice := model.Choice{
		ID:         "1",
		QuestionID: input.QuestionID,
		ChoiceText: input.ChoiceText,
	}
	return repository.CreateChoice(choice)
}
