package test

import(
	"testing"
	"context"
	"questionsanswers/pkg/model"
	"questionsanswers/pkg/data"
	"fmt"
	"reflect"

)


func TestCreateQuestion(t *testing.T){

	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
        { 1, 2, []model.Question{model.Question{ID:1,Subject:"",Description:""} }, []model.Question{ ID:2,model.Question{Subject:"",Description:""}}},
		{ 3, 4,[]model.Question{model.Question{ID:3,Subject:"A",Description:"B"} }, []model.Question{ID:4, model.Question{Subject:"A",Description:"B"}}},
    }

	for _, tt := range tests{
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
		
			db:= data.GetMongoDB()
			ctx := context.Background()
			
								
			mongodb ,err:=data.NewRepo(db,nil)
		
			if err !=nil{
				fmt.Println("Error occured inside CreateQuestion in repo")
			}
			for index, zz := range tt.questions{
			
				result := mongodb.CreateQuestion(ctx,zz)

				if result == nil {
					results,err:= mongodb.GetQuestionByID(ctx,int(zz.ID))
					if err!=nil{
						
						fmt.Println("GetQuestionByID mongo repo",  tt.questions, tt.want)
					}
					if !(reflect.DeepEqual(results, tt.want[index])) {
						t.Errorf("got %s, want %s", tt.questions, tt.want)
					}else{
						t.Logf("Success OK!")
					}


				}

			}
			
            
			
        })
    }  

}
/*

func TestGetAllQuestions(t *testing.T ){
	

	results := data.GetAllQuestion()	

	var want = []model.Question{
		model.Question {
			IDQuestion: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},
		model.Question {
			IDQuestion: 2, 
			Subject:"Maths 2", 
			Description: "Algorithms 2", 
		},
	}


	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}


func TestGetAllQuestionsByUser(t *testing.T) {
	results := data.GetAllQuestionsByUserMongo(1)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}


func TestGetAllQuestionsByUser(t *testing.T) {
	results := data.CreateQuestionMongo()	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}


func TestDeleteQuestion(t *testing.T) {
	results := data.DeleteQuestionMongo(2)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithms",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}

func TestGetUpdateQuestion(t *testing.T) {

	answer:= Answer{			
			IDQuestion : 1 ,
			Description : "Algorithm 3", 		
	}

   
	results := data.UpdateQuestionMongo(answer)	

	var want = []Question{
		Question {
			IdUser: 1, 
			Subject:"Maths", 
			Description: "Algorithm 3",
		},		
	}

	for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
            if !Equal(results, tt.want) {
                t.Errorf("got %d, want %d", results, tt.want)
            }else{
				t.Logf("Success OK!")
			}
			
        })
    }  
	
}

*/