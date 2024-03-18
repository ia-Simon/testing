package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"go_gql_tests_user_microservice/database"
)

type Resolver struct {
	DatabaseService *database.DatabaseService
}
