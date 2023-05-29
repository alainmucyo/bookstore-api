package main

import (
	"bookstore/core/environment"
	"bookstore/core/users"
	"bookstore/handlers/middleware"
	users2 "bookstore/handlers/users"
	hash2 "bookstore/store/hash"
	jwt2 "bookstore/store/jwt"
	"bookstore/store/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"
	"net/http"
	"os"
	"runtime"
)

func router(usersHandler *users2.Handler, a *middleware.AuthMiddleware) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/api-doc", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong pong")
	})

	r.POST("/login", usersHandler.Login)
	r.POST("/register", usersHandler.Register)
	s := r.Use(a.Authorize())

	{
		// users
		s.GET("/check", usersHandler.Check)
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

	authMiddleware := middleware.New(jwt)

	r := router(usersHandler, authMiddleware)
	err = r.Run(":" + envs.Port)
	if err != nil {
		log.Fatal(err)
	}
}
