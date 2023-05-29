package main

import (
	"bookstore/core/authors"
	"bookstore/core/books"
	"bookstore/core/environment"
	"bookstore/core/users"
	authors2 "bookstore/handlers/authors"
	books2 "bookstore/handlers/books"
	"bookstore/handlers/middleware"
	users2 "bookstore/handlers/users"
	hash2 "bookstore/store/hash"
	jwt2 "bookstore/store/jwt"
	"bookstore/store/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"
	"runtime"
)

func router(usersHandler *users2.Handler, authorsHandler *authors2.Handler, booksHandler *books2.Handler, a *middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Create a group for /api/v1 routes
	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong pong")
	})

	v1.POST("/login", usersHandler.Login)
	v1.POST("/register", usersHandler.Register)

	v1.GET("/books", booksHandler.FindAll)
	v1.GET("/books/:id", booksHandler.FindById)

	s := v1.Use(a.Authorize())

	{
		s.GET("/check", usersHandler.Check)

		// Authors routes
		s.GET("/authors", authorsHandler.FindAll)
		s.GET("/authors/:id", authorsHandler.FindById)
		s.POST("/authors", authorsHandler.Store)
		s.PUT("/authors/:id", authorsHandler.Update)
		s.DELETE("/authors/:id", authorsHandler.Delete)

		// Books routes
		s.POST("/books", booksHandler.Store)
		s.PUT("/books/:id", booksHandler.Update)
		s.DELETE("/books/:id", booksHandler.Delete)
	}

	return r
}

func main() {
	err := godotenv.Load()
	cpus := runtime.NumCPU()
	println(cpus)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envs := environment.New(os.Getenv("PORT"), os.Getenv("DB_URL"), os.Getenv("JWT_SECRET"))

	db := postgres.New(envs)
	db.Migrate()
	hash := hash2.New()

	jwt := jwt2.New(envs)

	userService := users.New(db, hash)
	usersHandler := users2.New(userService, jwt)

	authorService := authors.New(db)
	authorHandler := authors2.New(authorService)

	bookService := books.New(db)
	bookHandler := books2.New(bookService)

	authMiddleware := middleware.New(jwt)

	r := router(usersHandler, authorHandler, bookHandler, authMiddleware)
	err = r.Run(":" + envs.Port)
	if err != nil {
		log.Fatal(err)
	}
}
