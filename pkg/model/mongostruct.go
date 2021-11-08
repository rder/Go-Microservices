package model

type QuestionMongo struct {
	IDUser 		int     `bson:"id_user"`
	Subject    	string	`bson:"subject"`
	Description  string  `bson:"description"`
}

type AnswerMongo struct {
	IDQuestion int `bson:"id_question"`
	Description  string  `bson:"description"`
}

