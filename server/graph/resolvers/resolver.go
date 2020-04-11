package graph

import "goscrum/goscrum/server/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userService services.UserService
}

func NewResolver(userService services.UserService) Resolver {
	return Resolver{userService: userService}
}
