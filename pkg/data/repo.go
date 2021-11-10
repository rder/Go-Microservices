package data

import (
	"context"
	"database/sql"
	"encoding/json"
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
		UpdateQuestion (ctx context.Context,answer model.Question) (int, error)
		DeleteQuestion (ctx context.Context,answer model.Question) (int, error)

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

		err := repo.db.C("Questions").Insert(question)
		if err != nil {
			fmt.Println("Error occured inside CreateQuestion in repo",err)
			return err
		} else {
			fmt.Println("Question Created:", question.ID)
		}
		return nil
	}

	func (repo *mongorepo) GetAllQuestion(ctx context.Context) ([]model.Question, error) {
		fmt.Println("All question mongo repo", repo.db)
		
		var questions []bson.M
		
		err := repo.db.C(QuestionCollection).Find(nil).All(&questions)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Results All: ", questions) 
		}
				
	
		var q []model.Question

		rawjson, err1 := json.Marshal(&questions)
		json.Unmarshal(rawjson, &q)
		if err1!=nil{
			fmt.Println(" GetAllQuestion RESULT repo", err1)

		}
		return q, nil
		
	}



	func (repo *mongorepo) GetAllQuestionsByUser(ctx context.Context, idUser int) ([]model.Question, error) {
		fmt.Println("GetAllQuestionsByUser repo", repo.db)
		
		var questions []bson.M		

		err := repo.db.C(QuestionCollection).Find(bson.M{"idUser":idUser}).All(&questions)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Results All: ", questions) 
		}
			
		var q []model.Question

		rawjson, err1 := json.Marshal(&questions)
		json.Unmarshal(rawjson, &q)
		if err1!=nil{
			fmt.Println(" GetAllQuestion RESULT repo", err1)

		}
		return q, nil
	}

	


	func (repo *mongorepo) GetQuestionByID(ctx context.Context, id int) (model.Question, error) {
		fmt.Println(" GetQuestionByID mongo repo", repo.db)
		
		var question interface{}
		result:= repo.db.C(QuestionCollection).Find(bson.M{"_id": id}).One(&question)
		
		if result!=nil {
			fmt.Println("NO QUESTION : ",result)
		}
		var q model.Question

		rawjson, err1 := json.Marshal(question)
		json.Unmarshal(rawjson, &q)
		if err1!=nil{
			fmt.Println(" GetQuestionByID RESULT repo", err1)

		}
		
		fmt.Println(" GetQuestionByID RESULT repo", q)
		return q, nil
	}


	func (repo *mongorepo) UpdateQuestion(ctx context.Context, answer model.Question) (int, error) {
		
		fmt.Println(" UpdateQuestion", repo.db)
		
		var ans interface{}

		result:= repo.db.C(QuestionCollection).Find(bson.M{"_id": answer.ID}).One(&ans)
		
		if result!=nil { //			
			colQuerier := bson.M{"_id": answer.ID}
			change := bson.M{"$set": bson.M{"answer": answer.Answer}}
			err := repo.db.C(QuestionCollection).Update(colQuerier, change)
			if err != nil {
				fmt.Println("ERROR UpdateQuestion", err)
				panic(err)
			}			
		}
		fmt.Println("RESULT UpdateQuestion", result)
		return int(answer.ID), nil				
	}


	func (repo *mongorepo) DeleteQuestion(ctx context.Context, question model.Question) (int, error) {
		fmt.Println("DeleteQuestion mongo Question repo", question)
				
		errQ := repo.db.C(QuestionCollection).Remove(bson.M{"_id": question.ID})
		if errQ != nil {
			fmt.Println("ERROR DeleteQuestion", errQ)
		}
		
		return int(question.ID) ,errQ
	}

