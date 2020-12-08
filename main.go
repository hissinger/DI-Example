package main

import (
	"main/controller"
	"main/repository"

	"go.uber.org/dig"

	"fmt"
)

func newUserController(repo repository.Repository) controller.Controller {
	return &controller.User{Repository: repo}
}

func newMysqlRepository() repository.Repository {
	return &repository.Mysql{}
}

func main() {
	container := dig.New()
	container.Provide(newMysqlRepository)
	container.Provide(newUserController)

	container.Invoke(func(c controller.Controller) {
		fmt.Println(c.GetName())
	})
}
