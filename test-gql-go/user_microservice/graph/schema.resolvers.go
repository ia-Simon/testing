package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"go_gql_tests_user_microservice/database"
	"go_gql_tests_user_microservice/graph/model"

	"github.com/sirupsen/logrus"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.UserMutationPayload, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "Mutation#createUser",
		"input":  input,
	})
	logger.Info("Create user")
	var response model.UserMutationPayload

	userEntity, err := r.Resolver.DatabaseService.InsertUser(ctx, database.InsertUserInput{
		Name: input.Name,
		Age:  input.Age,
	})
	if err != nil {
		response.Status = false
		response.Errors = append(response.Errors, &model.MutationPayloadError{
			Code:    "USR00001",
			Message: "Falha ao criar usuário.",
		})
		logger.WithField("response", response).WithError(err).Error("Failed to create user.")
		return &response, nil
	}

	response.Status = true
	response.Resource = &model.User{
		ID:   userEntity.Id,
		Name: userEntity.Name,
		Age:  userEntity.Age,
	}

	logger.WithField("response", response).Info("User created successfully")
	return &response, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "Query#user",
		"id":     id,
	})
	logger.Info("Find user by id")

	userEntity, err := r.Resolver.DatabaseService.FindUserById(ctx, id)
	if err != nil {
		logger.WithError(err).Error("Failed to retrieve user")
		return nil, nil
	}

	if userEntity == nil {
		logger.Info("User not found")
		return nil, nil
	}

	user := &model.User{
		ID:   userEntity.Id,
		Name: userEntity.Name,
		Age:  userEntity.Age,
	}

	logger.WithField("response", user).Info("User found successfully")
	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, page int, size int) (*model.UserPagination, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "Query#users",
		"page":   page,
		"size":   size,
	})
	logger.Info("List users")
	var response model.UserPagination

	response.Page = page
	response.Size = size

	userEntities, total, err := r.Resolver.DatabaseService.FindUsers(ctx, page, size)
	if err != nil {
		logger.WithError(err).Error("Failed to retrieve users")
		return &response, err
	}

	for _, userEntity := range userEntities {
		response.Items = append(response.Items, &model.User{
			ID:   userEntity.Id,
			Name: userEntity.Name,
			Age:  userEntity.Age,
		})
	}
	response.Total = total

	logger.WithField("response", response).Info("Users retrieved successfully")
	return &response, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }