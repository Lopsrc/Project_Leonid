package main

import (
	"context"
	"fmt"

	// "go.mod/internal/authdata"
	// "go.mod/internal/authdata/db"
	"go.mod/internal/userdata"
	"go.mod/internal/userdata/db"
	"go.mod/internal/config"
	"go.mod/pkg/client/postgresql"
	"go.mod/pkg/logging"
)
//Решить проблему в userdata postgresql.go create() queryrow() не возвращает no rows in result set
//проверить findone()


func main(){
	// auth := authdata.AuthData{
	// 	Login: "IvanPupkin2002@mail.ru",
	// 	State: true,
	// 	Access_token: "access_token_02",
	// 	Refresh_token: "refresh_token_02",
	// }
	data := userdata.UserData{
		Name: "Ivan",
		Sex: "man",
		Birthdate: "2002-08-06",
		Weight: 88,
	}
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	client , err :=postgresql.NewClient(context.TODO(),3, cfg.Storage)
	if err!=nil{logger.Fatal("%v", err)}
	// repository := authdb.NewRepository(client, logger)
	repo := userdb.NewRepository(client, logger)
	err = repo.Create(context.TODO(), &data, 3)
	if err!=nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println(data.Id)
	// err = repository.Create(context.TODO(), &auth)
	// if err!=nil {logger.Fatalf("%v", err)}
	// fuser, err :=repository.FindOne(context.TODO(), "serpan2002@mail.ru")
	// if err!=nil{logger.Fatalf("%v", err)}
	fmt.Println("fuser")


}