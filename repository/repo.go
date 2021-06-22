package repository

import (
	"github.com/wambugucoder/golang-graphql/config"
	"github.com/wambugucoder/golang-graphql/graph/model"
)

// CreateQuestion -> DB call To Create Question
func CreateQuestion(question model.Question) (*model.Question, error) {
	config.DB.Create(&question)
	return &question, nil
}

func CreateChoice(choice model.Choice) (*model.Choice, error) {
	question := model.Question{}
	config.DB.First(&question, choice.QuestionID)
	choice.Question = &question
	err := config.DB.Create(&choice).Error
	if err != nil {
		return nil, err
	}
	return &choice, nil
}

func GetAllQuestions() ([]*model.Question, error) {
	var question []*model.Question
	config.DB.Find(&question)
	return question, nil
}
func GetAllChoices() ([]*model.Choice, error) {
	var choices []*model.Choice
	config.DB.Find(&choices)
	return choices, nil
}
