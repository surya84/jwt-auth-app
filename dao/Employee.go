package dao

import (
	"context"
	custommodels "jwtgen/customModels"
	"jwtgen/graph/model"
	"jwtgen/initializers"

	"github.com/rs/zerolog/log"
)

type Emp interface {
	CreateUser(ctx context.Context, input model.EmpData) (string, error)
}

func CreateUser(ctx context.Context, input model.EmpData) (string, error) {
	employee := custommodels.Employee{
		Name:  input.Name,
		Email: input.Email,
		EmpID: input.EmpID,
	}
	err := initializers.DB.Create(&employee)
	if err.Error != nil {
		log.Printf("error in creating user :: ", err.Error)
		return "", err.Error
	}
	return "User created", nil
}

func GetEmpData(ctx context.Context, input model.UserRequest) (*model.Employee, error) {
	var emp custommodels.Employee
	err := initializers.DB.Where("email=?", input.Email).First(&emp)
	if err.Error != nil {
		log.Printf("error in fetching user details :: ", err.Error)
		return &model.Employee{}, err.Error
	}
	return &model.Employee{
		Name:  emp.Name,
		EmpID: emp.EmpID,
		Email: emp.Email,
	}, nil
}

func AuthorizeUser(email string) (*custommodels.Employee, error) {
	var user custommodels.Employee

	err := initializers.DB.Where("email=?", email).First(&user)
	if err.Error != nil {
		log.Printf("error in getting data from db :: ", err)
		return &custommodels.Employee{}, err.Error
	}
	return &user, nil
}
