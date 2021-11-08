package data

import (
	"context"
	"database/sql"
	"fmt"
	"questionsanswers/pkg/model"

	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


	const QuestionCollection = "Questions"
	const AnswerCollection = "Answers"


	type Repository interface {
		CreateQuestion(ctx context.Context, question model.Question) error
		GetAllQuestion(ctx context.Context) ([]model.Question, error)
		GetAllQuestionsByUser (ctx context.Context,id int) ([]model.Question , error)
		GetQuestionByID (ctx context.Context,id int) (model.Question , error)
		UpdateQuestion (ctx context.Context,answer model.Answer) (int, error)
		DeleteQuestion (ctx context.Context,answer model.Answer) (int, error)

	}
		
	// Repo structure for mongodb
	type mongorepo struct {
		db     *mgo.Database
		logger log.Logger
	}

	type sqlrepo struct {
		sqldb  *sql.DB
		logger log.Logger
	}

	// Repo for mongodb
	func NewRepo(db *mgo.Database, logger log.Logger) (Repository, error) {
		return &mongorepo{
				db:db,
				logger: log.With(logger, "repo", "mongodb"),}, nil
	}


	/*All Methods*/

	func (repo *mongorepo) CreateQuestion(ctx context.Context, question model.Question) error {
		fmt.Println("create question mongo repo", repo.db)
		err := repo.db.C(QuestionCollection).Insert(ctx,question)
		if err != nil {
			fmt.Println("Error occured inside CreateQuestion in repo")
		return err
		} else {
			fmt.Println("Question Created:", question.ID)
		}
		return nil
	}

	func (repo *mongorepo) GetAllQuestion(ctx context.Context) ([]model.Question, error) {
		fmt.Println("All question mongo repo", repo.db)
		
		var questions []model.Question		

		err := repo.db.C(QuestionCollection).Find(ctx).All(&questions)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Results All: ", questions) 
		}
			
		return questions, nil
		
	}



	func (repo *mongorepo) GetAllQuestionsByUser(ctx context.Context, idUser int) ([]model.Question, error) {
		fmt.Println("GetAllQuestionsByUser repo", repo.db)
		
		var questions []model.Question		

		err := repo.db.C(QuestionCollection).Find(bson.M{"_id":idUser}).All(&questions)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Results All: ", questions) 
		}
			
		return questions, nil
	}

	


	func (repo *mongorepo) GetQuestionByID(ctx context.Context, id int) (model.Question, error) {
		fmt.Println(" GetQuestionByID mongo repo", repo.db)
		
		var question model.Question
		result:= repo.db.C(QuestionCollection).Find(bson.M{"_id": id}).All(&question)
		
		if result!=nil {
			fmt.Println("NO QUESTION : ")
			panic(result)
		}

		return question, nil
	}


	func (repo *mongorepo) UpdateQuestion(ctx context.Context, answer model.Answer) (int, error) {
		
		var ans model.Answer
		result:= repo.db.C(AnswerCollection).Find(bson.M{"id_question": answer.IDQuestion}).One(&ans)
		
		if result!=nil { // No ANSWER INSERTED BEFORE
			err1 := repo.db.C(AnswerCollection).Insert(ctx,answer)
			if err1 != nil{
				panic(err1)
			}

		}else{		
			colQuerier := bson.M{"id_question": answer.IDQuestion}
			change := bson.M{"$set": bson.M{"description": answer.Description}}
			err := repo.db.C(AnswerCollection).Update(colQuerier, change)
			if err != nil {
				panic(err)
			}
	}
		return int(answer.IDQuestion), nil
				
	}


	func (repo *mongorepo) DeleteQuestion(ctx context.Context, answer model.Answer) (int, error) {
		fmt.Println("DeleteQuestion mongo repo", repo.db)
		
		err := repo.db.C(AnswerCollection).Remove(bson.M{"_id": &answer.IDAnswer})
		if err != nil {
			panic(err)
			
		}
		errQ := repo.db.C(QuestionCollection).Remove(bson.M{"_id": &answer.IDAnswer})
		if errQ != nil {
			panic(errQ)
		}
		
		return int(answer.IDQuestion) ,err
	}

