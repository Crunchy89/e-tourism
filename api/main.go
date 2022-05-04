package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_userHandler "api/service/user/handler"
	_userRepo "api/service/user/repository"
	_userUseCase "api/service/user/usecase"
	"api/utils/s"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-contrib/cors"
)

func main() {
	println("Initialize server")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "etourism"
	}

	// Setup the Repository
	mongoAuthSource := os.Getenv("DB_AUTH_SOURCE")
	if mongoAuthSource == "" {
		mongoAuthSource = "localhost:27017"
	}
	userName := os.Getenv("DB_USER_NAME")
	if userName == "" {
		userName = "etourism"
	}
	passWord := os.Getenv("DB_PASSWORD")
	if passWord == "" {
		passWord = "development"
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", userName, passWord, mongoAuthSource, databaseName)
	mongoDBInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	database := mongoDBInstance.Database(databaseName)
	// mongoClient := mg.NewMongoClient(mongoDBInstance)

	// setup the repository
	userRepo := _userRepo.NewUserRepository(database)

	// Setup the usecase
	userUseCase := &_userUseCase.UserUseCase{
		User: userRepo,
		// Client:        mongoClient,
	}

	// Setup the handler
	userHandler := &_userHandler.UserHandler{
		User: userUseCase,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	}))
	v1 := r.Group("api").Group("v1")
	user := v1.Group("user")
	user.GET("username", userHandler.FetchUserByUsername)
	user.GET("token", userHandler.FetchUserByToken)
	user.GET("id", userHandler.FetchUserById)
	user.POST("", userHandler.AddUser)

	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to API e-tourism Lombok Tengah")
	})
	r.GET("ping", func(c *gin.Context) {
		if err := mongoDBInstance.Ping(c, readpref.Primary()); err != nil {
			s.Abort(c, err)
			return
		}
		c.JSON(200, gin.H{
			"server":  "OK",
			"mongodb": "OK",
			"message": "pong",
		})
	})
	// Put port just for testing another test again
	port := ":" + os.Getenv("PORT")
	if port == ":" || port == "" {
		port = ":" + "8080"
	}
	r.Run(port)
}
