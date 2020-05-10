package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"goscrum/goscrum/server/constants"
	"goscrum/goscrum/server/di"
	"goscrum/goscrum/server/graph/gqlgen_gen"
	"goscrum/goscrum/server/validator"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin" //nolint
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ginLambda *ginadapter.GinLambda

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	validator.Add()
	resolver := di.InitializeResolver(os.Getenv(constants.Debug) == "1")
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(gqlgen_gen.NewExecutableSchema(gqlgen_gen.Config{Resolvers: &resolver}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// func cognitoUserToContext() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		apiGwContext, _ := ginLambda.GetAPIGatewayContext(c.Request)
//		spew.Dump(apiGwContext)
//		// c.Request = c.Request.WithContext(ctx) TODO
//		c.Next()
//	}
//}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

//nolint
func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	//r.Use(cognitoUserToContext())
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": time.Now(),
		})
	})

	ginLambda = ginadapter.New(r)
}

//nolint
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
