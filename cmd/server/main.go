// cmd/server/main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" // Você pode escolher uma biblioteca de roteamento, como o Gorilla Mux.
	//"nba_api/internal/api/handlers"
	//"nba_api/internal/api/routes"
	// "nba_api/config"
)

func main() {
	// Carregar configurações do arquivo .env
	//config.LoadEnv()

	// Inicializar o roteador
	router := mux.NewRouter()

	// Conectar ao banco de dados (se necessário)
	// db := database.ConnectDB()
	// defer db.Close()

	// Inicializar os controladores e rotas
	//h := handlers.NewHandler()
	//routes.SetupRoutes(router, h)

	// Iniciar o servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão, você pode alterá-la no arquivo .env
	}

	log.Printf("Server started on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
