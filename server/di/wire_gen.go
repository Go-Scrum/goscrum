// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"goscrum/goscrum/server/aws"
	"goscrum/goscrum/server/graph/resolvers"
	"goscrum/goscrum/server/services"
)

// Injectors from wire.go:

func InitializeResolver(debug bool) graph.Resolver {
	db := aws.NewDBClient(debug)
	userService := services.NewUserService(db)
	resolver := graph.NewResolver(userService)
	return resolver
}

func InitializeUserService(debug bool) services.UserService {
	db := aws.NewDBClient(debug)
	userService := services.NewUserService(db)
	return userService
}
