package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"go_gql_tests_user_microservice/graph/model"

	"github.com/sirupsen/logrus"
)

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"action": "Entity#findUserByID",
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

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }