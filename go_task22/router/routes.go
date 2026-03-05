package router

import (
	"gorm/handlers"
	"gorm/repository"
	"gorm/service"
	"net/http"

	"gorm.io/gorm"
)


func Routes(db *gorm.DB) {
	repo:=repository.New(db)
	services:=service.New(repo)
	handler:=handlers.New(services)
	http.HandleFunc("/create", handler.CreateUser)
	http.HandleFunc("/get",handler.GetUser)
	http.HandleFunc("/getbyid/{id}",handler.GetUserById)
	http.HandleFunc("/updatebyid/{id}",handler.UpdateUserById)
	http.HandleFunc("/patchbyid/{id}",handler.PatchByid)
	http.HandleFunc("/deletebyid/{id}",handler.DeleteUserById)
}
