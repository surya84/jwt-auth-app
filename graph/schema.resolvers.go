package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"jwtgen/dao"
	"jwtgen/graph/model"
	grpcclient "jwtgen/grpcClient"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.EmpData) (string, error) {
	data, err := dao.CreateUser(ctx, input)
	if err != nil {
		return "", err
	}
	return data, nil
}

// GetUserData is the resolver for the GetUserData field.
func (r *queryResolver) GetUserData(ctx context.Context, input model.UserRequest) (*model.Employee, error) {

	client := grpcclient.GrpcClinet()
	resp, err := grpcclient.GetData(client, input.Email)

	if err != nil {
		return nil, err
	}
	return &model.Employee{
		Name:  resp.Name,
		Email: resp.Email,
		EmpID: resp.EmpID,
	}, nil

	// EmpData, err := dao.GetEmpData(ctx, input)
	// if err != nil {
	// 	return nil, err
	// }
	// return EmpData, nil
}

// Hello is the resolver for the Hello field.
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello world Hi ", nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
