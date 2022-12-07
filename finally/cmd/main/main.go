package main

import (
	"context"
	"fmt"

	// "go.mod/internal/authdata"
	// "go.mod/internal/authdata/db"
	//"go.mod/internal/authdata"
	"go.mod/internal/config"
	"go.mod/internal/userdata"
	"go.mod/internal/userdata/db"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
)

//Решить проблему в userdata postgresql.go create() queryrow() не возвращает no rows in result set возможно scan() должен считывать string 
//проверить findone()
// const(
// 	CREATE = 1
// 	FIND = 2
	// DELETE = 3
	// CHANGE = 4
// )
// type Repos struct{
	// repositoryauth authdata.Repository
	// repositorydata userdata.Repository
	// 
// }
// func manage(action int){
// 
	// switch action {
	// case CREATE:
	// case FIND:
	// case DELETE:
	// case CHANGE:
// 
		// 
	// }
// }
// 
func main(){
	// auth := authdata.AuthData{
	// 	Login: "IvanPupkin2002@mail.ru",
	// 	State: true,
	// 	Access_token: "access_token_02",
	// 	Refresh_token: "refresh_token_02",
	// }
	data := userdata.UserData{
		Id: 2,
		Name: "Ivan",
		Sex: "man",
		Birthdate: "2001-02-19",
		Weight: 90,
	}
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	client , err :=postgresql.NewClient(context.TODO(),3, cfg.Storage)
	if err!=nil{logger.Fatal("%v", err)}
	// repository := authdb.NewRepository(client, logger)


	repo := userdb.NewRepository(client, logger)
	// err = repo.Create(context.TODO(),&data, 2 )
	err = repo.Update(context.TODO(), &data)
	// err = repo.Delete(context.TODO(), "2")
	if err!=nil {
		logger.Fatalf("%v", err)
	}
	//fmt.Println(data.Id)


	// err = repository.Create(context.TODO(), &auth)
	// if err!=nil {logger.Fatalf("%v", err)}
	// fuser, err :=repository.FindOne(context.TODO(), "serpan2002@mail.ru")
	// if err!=nil{logger.Fatalf("%v", err)}
	fmt.Println("fuser")

	// manage(CREATE)
}