package model



type Question struct {
    ID    int32   `json:"id" bson:"_id"`
    IDUser   int32  `json:"idUser" bson:"idUser"`
    Subject string `json:"subject" bson:"subject"`
    Description  string `json:"description" bson:"description"`
    Answer  string `json:"answer" bson:"answer"`
}