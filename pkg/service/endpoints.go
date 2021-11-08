package service

import (
	"context"
	"questionsanswers/pkg/model"

	"github.com/go-kit/kit/endpoint"
)


type Endpoints struct{
	GetAllQuestions endpoint.Endpoint
    GetAllQuestionsByUser endpoint.Endpoint
	CreateQuestion endpoint.Endpoint
	GetQuestionByID endpoint.Endpoint
	UpdateQuestion endpoint.Endpoint
	DeleteQuestion endpoint.Endpoint
}


func MakeEndpoints (s Service) Endpoints{
	return Endpoints{
		GetAllQuestions: makeGetAllQuestionsEndpoint(s),
		GetAllQuestionsByUser: makeGetAllQuestionsByUserEndpoint(s),
		CreateQuestion : makeCreateQuestionEndpoint(s),		
		GetQuestionByID : makeGetQuestionByIDEndpoint(s),	
    	UpdateQuestion : makeUpdateQuestionEndpoint(s),
    	DeleteQuestion : makeDeleteQuestionEndpoint(s),

	}
}

func makeGetAllQuestionsEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,request interface{})(interface{},error)	{
		
		question,_:= s.GetAllQuestions(ctx)
		return question,nil
	}
	
}

func makeGetAllQuestionsByUserEndpoint(s Service) endpoint.Endpoint{

	return func(ctx context.Context,request interface{})(interface{},error)	{
		req := request.(model.CreateQuestionRequest)
		question,_:= s.GetAllQuestionsByUser(ctx,int(req.IDUser))

		return question,nil
	}

	
}

func makeCreateQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error)	{
		req := request.(model.CreateQuestionRequest)
		q,err:= s.CreateQuestion(ctx,model.Question{Subject:req.Subject,Description:req.Description,IDUser:req.IDUser})
		return q,err
	}
}

func makeGetQuestionByIDEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{})(interface{},error){
		req := request.(model.GetQuestionByIDRequest)
		q,_:= s.GetQuestionByID(ctx,int(req.ID))
		return q,nil
	}
}

func makeUpdateQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(model.CreateAnswerRequest)
		q,_:= s.UpdateQuestion(ctx,model.Answer{IDQuestion:req.IDquestion,Description:req.Description})
		return q,nil
	}

}

func makeDeleteQuestionEndpoint (s Service) endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req := request.(model.CreateAnswerRequest)
		q,_:= s.DeleteQuestion(ctx,model.Answer{IDQuestion:req.IDquestion,Description:req.Description})

		return q,nil
	}
}

    	