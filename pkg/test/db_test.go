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
        { 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"} }, []model.Question{ model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"}}},

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
			
				fmt.Println("Question",  zz)
				result := mongodb.CreateQuestion(ctx,zz)

				if result == nil {
					
					if !(reflect.DeepEqual(zz, tt.want[index])) {
						fmt.Println("NOT EQUAL mongo repo",  tt.questions, tt.want)
					}else{
						t.Logf("Success OK!")
					}


				}

			}


        })
    }
}

func TestGetQuestionByID(t *testing.T){

	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
		{ 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"} }, []model.Question{ model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"}}},
	}

	for _, tt := range tests{
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T){      
		
			db:= data.GetMongoDB()
			ctx := context.Background()			
								
			mongodb ,err:=data.NewRepo(db,nil)
		
			if err !=nil{
				fmt.Println("Error occured inside GetQuestionByID in repo")
			}
			for index, zz := range tt.questions{
			
				fmt.Println("Question",  zz)
				result,err := mongodb.GetQuestionByID(ctx,int(zz.ID))
				
				if err != nil{						
						fmt.Println("GetQuestionByID ERROR repo",  tt.questions, tt.want)
					}					
					if !(reflect.DeepEqual(result, tt.want[index])) {
						fmt.Println("NOT EQUAL mongo repo", result, tt.want)
					}else{
						t.Logf("Success OK!")
					}
				}
        })
	}
}


func TestGetAllQuestion(t *testing.T){
	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
		{ 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"} }, []model.Question{ model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"}}},
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
			
			result,err := mongodb.GetAllQuestion(ctx)

			if err != nil {
				fmt.Println("ERROR",  err)
	
			}		
			if !(reflect.DeepEqual(result, tt.want)) {
				fmt.Println("NOT EQUAL mongo repo",  result, tt.want)
			}else{
				t.Logf("Success OK!")
			}			

		})				
    
    }  
			
}


func TestGetAllQuestionsByUser(t *testing.T){
	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
		{ 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"} }, []model.Question{ model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD"}}},
    }


	for _, tt := range tests{
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
		
			db:= data.GetMongoDB()
			ctx := context.Background()
			
								
			mongodb ,err:=data.NewRepo(db,nil)
		
			if err !=nil{
				fmt.Println("Error GetAllQuestionsByUser in repo",err)
			}
			
			result,err := mongodb.GetAllQuestionsByUser(ctx,2)

			if err != nil {
				fmt.Println("ERROR",  err)
	
			}		
			if !(reflect.DeepEqual(result, tt.want)) {
				fmt.Println("NOT EQUAL mongo repo",  result, tt.want)
			}else{
				t.Logf("Success OK!")
			}			

		})				
    
    }  
			
}


func TestUpdateQuestion(t *testing.T){
	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
		{ 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD",Answer:"JH"} }, []model.Question{ model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD",Answer:"JH"}}},
		
	}

	for _, tt := range tests{
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
		
			db:= data.GetMongoDB()
			ctx := context.Background()
			
								
			mongodb ,err:=data.NewRepo(db,nil)
		
			if err !=nil{
				fmt.Println("Error occured inside UpdateQuestion in repo")
			}
		

			for index, zz := range tt.questions{
			
				fmt.Println("ANSWER",  zz)
				result,err := mongodb.UpdateQuestion(ctx,zz)
				resultI,errI := mongodb.GetQuestionByID(ctx,int(zz.ID))

				if err == nil && errI!=nil{	

					if !(reflect.DeepEqual(resultI, tt.want[index])) {
						fmt.Println("UPDATE",  result, tt.want)
						fmt.Println("NOT EQUAL mongo repo",  resultI, tt.want)
					}else{
						t.Logf("Success OK!")
					}
				}
			}	

		})			
    
    }  
			
}




func TestDeleteQuestion(t *testing.T) {
	var tests = []struct {
        a, b int
		questions []model.Question
        want []model.Question
    }{
     	{ 1, 2, []model.Question{model.Question{ID:13,IDUser:1,Subject:"AA",Description:"DD",Answer:"JH"} }, []model.Question{} },
		
	
	
	}


	for _, tt := range tests{
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {      
		
			db:= data.GetMongoDB()
			ctx := context.Background()
			
								
			mongodb ,err:=data.NewRepo(db,nil)
		
			if err !=nil{
				fmt.Println("Error GetAllQuestionsByUser in repo",err)
			}
			
			result,err := mongodb.GetQuestionByID(ctx,13)
			resultD,errD:= mongodb.DeleteQuestion(ctx,result)

			if errD != nil {
				fmt.Println("ERROR",  err)
	
			}		
			if errD==nil{
				fmt.Println("Error DeleteQuestion in repo",resultD)	
			}else{
				t.Logf("Success OK!")
			}	
			//Last Test remove the DB just in case to restart the tests
			db.C("Questions").RemoveAll(nil)
		})    
    }  
}



