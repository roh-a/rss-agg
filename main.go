package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/roh-a/rss-agg/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found %v", err)
	}

	//Initialize database manager with Vault Integration
	var err error 
	dbManager, err := NewDatabaseManager()
	if err != nil {
		log.Fatalf("failed to initialize database manager: %v", err)
	}

	defer dbManager.Close()






	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environment")
	}




	apiConfig := apiConfig{
		DB : database.New(dbManager.GetDB()), 
	}



	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:	300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthy", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users", apiConfig.handleGetUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Port:", portString)
}
