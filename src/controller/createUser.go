package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wagner477/crud-golang/src/configuration/validation"
	"github.com/wagner477/crud-golang/src/controller/model/request"
	"github.com/wagner477/crud-golang/src/controller/model/respose"
)

func CreateUser(c *gin.Context) {
	log.Println("Create User Controller")

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to bind JSON %s", err.Error())

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)

		return
	}

	fmt.Println(userRequest)

	response := respose.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(200, response)

}
