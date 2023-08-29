package main

import (
	"context"
	"fmt"
	"log"
	"os"

	adminRoutes "github.com/Ashutosh-730/go-restro/routes/admin"
	userRoutes "github.com/Ashutosh-730/go-restro/routes/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// Load environment variables
	log.Println("---------:Loading environment variables:---------")
	// utils.LoadEnvVars()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	userRoutes.UserRoutes(user)

	admin := r.Group("/fund15/esthenos")
	adminRoutes.AdminRoutes(admin)

	return r
}

func LoadGinMiddleWares(router *gin.Engine) {
	// For local development, enable CORS middleware
	if os.Getenv("run_mode") == "local" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"} // Update with your frontend domain
		config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
		router.Use(cors.New(config))
	}
}

func main() {
	// Initialize Gin Lambda
	router := SetupRouter()

	// Load Gin middlewares
	LoadGinMiddleWares(router)

	runMode := os.Getenv("run_mode")
	switch runMode {
	case "lambda":
		// For AWS Lambda deployment, use the Gin Lambda adapter
		log.Println("Server running in lambda mode")
		// gin.SetMode(gin.ReleaseMode)
		ginLambda = ginadapter.New(router)

		// Start the Lambda handler
		lambda.Start(Handler)

	case "local":
		// For local development, enable Gin's recovery and logger middleware
		log.Println("Server running in local mode")
		port := os.Getenv("port")
		if port == "" {
			port = "2020"
		}
		fmt.Println("Server running on port:", port)
		log.Fatal(router.Run(":" + port))

	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}
