package model

type Question struct {
    ID        int32 `json:"id"`
    IDUser    int32 `json:"idUser"`
    Subject  string `json:"subject"`
    Description string `json:"description"`   
}

type Answer struct {
    IDAnswer        int32 `json:"idAnswer"`
    IDQuestion   int32 `json:"idQuestion"`
    Description string `json:"description"`   
}


