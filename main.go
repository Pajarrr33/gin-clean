package main

import (
	"gin-db/config"
	"gin-db/controller"
	"gin-db/repository"
	"gin-db/routes"
	"gin-db/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect into the database
	db := config.ConnectDb()

	defer db.Close()

	// Implement dependency injection
	var (
		// Repository
		credentialRepository repository.CredentialRepository = repository.NewCredentialRepo(db)

		// Usecase
		credentialUsecase usecase.CredentialUsecase = usecase.NewCredentialUsecase(credentialRepository)

		// Controller
		credentialController controller.CredentialController = controller.NewCredentialController(credentialUsecase)
	)

	server := gin.Default()

	// Routes
	routes.Credentials(server, credentialController)

	server.Run(":8080")
}
