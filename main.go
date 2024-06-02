package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/VSM1le/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in hte enviroment")
	}
	portDB := os.Getenv("DB_URL")
	if portDB == "" {
		log.Fatal("Port is not found in hte enviroment")
	}

	conn, err := sql.Open("postgres", portDB)
	if err != nil {
		log.Fatal("Can not connect to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerERR)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeed)
	v1Router.Post("/follow", apiCfg.middlewareAuth(apiCfg.handlerCreatFeedFollow))
	v1Router.Get("/follow", apiCfg.middlewareAuth(apiCfg.handlerSelectFeedFollow))
	v1Router.Delete("/follow/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))
	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlergetPostsForUser))

	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(apiCfg.DB, collectionConcurrency, collectionInterval)

	log.Printf("Server staring on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
