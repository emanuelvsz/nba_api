// cmd/server/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"nba_api/internal/api/handlers"
	"nba_api/internal/api/routes"
	"nba_api/config"
)

func main() {
	// Carregar configurações do arquivo .env
	config.LoadEnv()

	// Inicializar o roteador
	router := mux.NewRouter()

	// Conectar ao banco de dados (se necessário)
	// db := database.ConnectDB()
	// defer db.Close()

	// Inicializar os controladores e rotas
	handler := handlers.NewHandler()
	routes.SetupRoutes(router, handler)

	// Iniciar o servidor HTTP
	port := config.GetEnvVar("PORT")
	if port == "" {
		port = "8080" // Porta padrão, você pode alterá-la no arquivo .env
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
