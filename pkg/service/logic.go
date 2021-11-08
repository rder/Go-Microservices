package service

import (
	"context"
	"log"
	"questionsanswers/pkg/data"
	"questionsanswers/pkg/model"

	"github.com/go-kit/kit/log"
)


type srv struct{
	repository data.Repository
	logger log.Logger
}


func NewService (rep data.Repository, logger log.Logger) Service{

	return &srv{ 
		repository: rep,
		logger: logger,
	}
}

func (s srv) GetAllQuestions (ctx context.Context) ( []model.Question,error){
	
	return s.repository.GetAllQuestion(ctx)

}

func (s srv) GetAllQuestionsByUser (ctx context.Context, id int) ([]model.Question,error){
	
	return s.repository.GetAllQuestionsByUser(ctx,id)

}

func (s srv) CreateQuestion (ctx context.Context, question model.Question) (error){

	return s.repository.CreateQuestion(ctx,question)

}

func (s srv) UpdateQuestion (ctx context.Context, answer model.Answer) (int,error){

	return s.repository.UpdateQuestion(ctx,answer)

}

func (s srv) DeleteQuestion (ctx context.Context, answer model.Answer) (int,error){

	return s.repository.DeleteQuestion(ctx,answer)

}


func (s srv) GetQuestionByID (ctx context.Context,  id int) (model.Question,error){

	return s.repository.GetQuestionByID(ctx,id)

}