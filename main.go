package main

import (
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func dbConnect() *sql.DB {
    db, err := sql.Open("mysql", "user1:1234@tcp(127.0.0.1:3306)/test")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    return db
}

func getUserHandler(res http.ResponseWriter, req *http.Request) {
        var response string
        db := dbConnect()
	    row := db.QueryRow("SELECT * FROM User")
        err := row.Scan(&response)
        if err != nil && err != sql.ErrNoRows {
            log.Fatal(err)
        }
        res.Write([]byte(response))
}

func postUserHandler(res http.ResponseWriter, req *http.Request)  {}

func getSelectedUserHandler(res http.ResponseWriter, req *http.Request) {}
func patchSelecedtUserHandler(res http.ResponseWriter, req *http.Request) {}
func deleteSelecedtUserHandler(res http.ResponseWriter, req *http.Request) {}
func getSelecedtUserProductHandler(res http.ResponseWriter, req *http.Request) {}
func deleteSelectedUserProductHandler(res http.ResponseWriter, req *http.Request) {}

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Logger)

	router.Route("/user", func(router chi.Router) {
		router.Get("/", getUserHandler)
		router.Post("/", postUserHandler)
		router.Get("/{mail}", getSelectedUserHandler)
		router.Patch("/{mail}", patchSelecedtUserHandler)
		router.Delete("/{mail}", deleteSelecedtUserHandler)
		router.Get("/{mail}/product", getSelecedtUserProductHandler)
		router.Delete("/{mail}/product", deleteSelectedUserProductHandler)
	})

    http.ListenAndServe(":8080", router)
}
    