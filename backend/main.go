package main

import (
	"controle-de-gastos/pkg/config"
	"controle-de-gastos/pkg/database"
	"controle-de-gastos/src/handler/category_handler"
	"controle-de-gastos/src/handler/expense_handler"
	"controle-de-gastos/src/handler/incoming_handler"
	"controle-de-gastos/src/handler/user_handler"
	"controle-de-gastos/src/repository/category_repo"
	"controle-de-gastos/src/repository/expense_repo"
	"controle-de-gastos/src/repository/incoming_repo"
	"controle-de-gastos/src/repository/user_repo"
	"controle-de-gastos/src/routes"
	"controle-de-gastos/src/service/category_service"
	"controle-de-gastos/src/service/expense_service"
	"controle-de-gastos/src/service/incoming_service"
	"controle-de-gastos/src/service/user_service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
		panic(err)
	}

	// Carregar configurações do BD
	db_config := config.DbConfigLoad()

	// Realizar a conexao do BD
	dbConnection, err := database.ConnectDB(db_config)
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	router, api := routes.SetupRouter()

	// Injeção de dependencias
	routes.SetupUserRoutes(api, user_handler.NewUserHandler(user_service.NewUserService(user_repo.NewUsuarioRepo(dbConnection))))
	routes.SetupExpenseRoutes(api, expense_handler.NewExpenseHandler(expense_service.NewExpenseService(expense_repo.NewExpenseRepo(dbConnection))))
	routes.SetupCategoryRoutes(api, category_handler.NewCategoryHandler(category_service.NewCategoryService(category_repo.NewCategoryRepo(dbConnection))))
	routes.SetupIncomingRoutes(api, incoming_handler.NewIncomingHandler(incoming_service.NewIncomingService(incoming_repo.NewIncomingRepo(dbConnection))))

	// Iniciar servidor
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
