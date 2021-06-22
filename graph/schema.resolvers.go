package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/wambugucoder/golang-graphql/repository"
	"github.com/wambugucoder/golang-graphql/service"

	"github.com/wambugucoder/golang-graphql/graph/generated"
	"github.com/wambugucoder/golang-graphql/graph/model"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	return service.CreateQuestion(input)
}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
	return service.CreateChoice(input)
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	return repository.GetAllQuestions()
}

func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	return repository.GetAllChoices()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
