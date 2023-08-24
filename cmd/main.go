package main

import (
	"database/sql"
	"gabrielmaurici/first-go-api/internal/database"
	"gabrielmaurici/first-go-api/internal/routes"
	"gabrielmaurici/first-go-api/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gabras")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	postDb := database.NewPostDb(db)
	postUsecase := usecase.NewPostUseCase(postDb)
	postRoutes := routes.NewRouterPost(postUsecase)

	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Mount("/post", postRoutes.AddHandlerPost())

	http.ListenAndServe(":8080", r)
}
