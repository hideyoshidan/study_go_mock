package main

import (
	"fmt"

	"golang.org/x/exp/slog"
	"mock.com/repository"
	"mock.com/service"
)

func main() {
	ur := repository.NewUserRepository()
	us := service.NewUserService(ur)
	err := us.Create()
	if err != nil {
		slog.Error(fmt.Sprintf("Error has been ouccred while create user : %v", err))
	} else {
		slog.Info("It was successfull")
	}
}
