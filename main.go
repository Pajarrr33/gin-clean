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
		UserRepository repository.UserRepository = repository.NewUserRepo(db)

		// Usecase
		credentialUsecase usecase.CredentialUsecase = usecase.NewCredentialUsecase(credentialRepository)
		userUsecase usecase.UserUsecase = usecase.NewUserUsecase(UserRepository)

		// Controller
		credentialController controller.CredentialController = controller.NewCredentialController(credentialUsecase)
		userController controller.UserController = controller.NewUserController(userUsecase,credentialUsecase)
	)

	server := gin.Default()

	// Routes
	routes.Credentials(server, credentialController)
	routes.Users(server,userController)

	server.Run(":8080")
}
