//+build wireinject

package di

import (
	"goscrum/goscrum/server/aws"
	graph "goscrum/goscrum/server/graph/resolvers"
	"goscrum/goscrum/server/services"

	"github.com/google/wire"
)

func InitializeResolver(debug bool) graph.Resolver {
	wire.Build(graph.NewResolver, services.NewUserService, aws.NewDBClient)
	return graph.Resolver{}
}

func InitializeUserService(debug bool) services.UserService {
	wire.Build(services.NewUserService, aws.NewDBClient)
	return services.UserService{}
}
