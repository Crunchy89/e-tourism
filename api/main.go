package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	_fasilitasHandler "api/service/fasilitas/handler"
	_fasilitasRepo "api/service/fasilitas/repository"
	_fasilitasUseCase "api/service/fasilitas/usecase"
	_penginapanHandler "api/service/penginapan/handler"
	_penginapanRepo "api/service/penginapan/repository"
	_penginapanUseCase "api/service/penginapan/usecase"
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

	// Setup the Database
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		databaseName = "etourism"
	}

	mongoAuthSource := os.Getenv("DB_AUTH_SOURCE")
	if mongoAuthSource == "" {
		mongoAuthSource = "localhost:27017"
	}

	userName := os.Getenv("DB_USER_NAME")
	if userName == "" {
		userName = "root"
	}

	passWord := os.Getenv("DB_PASSWORD")
	if passWord == "" {
		passWord = "password"
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s?authSource=admin", userName, passWord, mongoAuthSource, databaseName)
	mongoDBInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	database := mongoDBInstance.Database(databaseName)

	// setup the repository
	userRepo := _userRepo.NewUserRepository(database)
	penginapanRepo := _penginapanRepo.NewPenginapanRepository(database)
	fasilitasRepo := _fasilitasRepo.NewFasilitasRepository(database)

	// Setup the usecase
	userUseCase := &_userUseCase.UserUseCase{
		User: userRepo,
	}

	penginapanUseCase := &_penginapanUseCase.PenginapanUseCase{
		Penginapan: penginapanRepo,
	}

	fasilitasUseCase := &_fasilitasUseCase.FasilitasUseCase{
		Fasilitas: fasilitasRepo,
	}

	// Setup the handler
	userHandler := &_userHandler.UserHandler{
		User: userUseCase,
	}

	penginapanHandler := &_penginapanHandler.PenginapanHandler{
		Penginapan: penginapanUseCase,
	}

	fasilitasHandler := &_fasilitasHandler.FasilitasHandler{
		Fasilitas: fasilitasUseCase,
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	}))

	v1 := r.Group("api").Group("v1")

	fasilitas := v1.Group("fasilitas")
	fasilitas.POST("", fasilitasHandler.AddFasilitas)
	fasilitas.GET("", fasilitasHandler.FetchFasilitas)
	fasilitas.PUT("", fasilitasHandler.UpdateFasilitas)
	fasilitas.DELETE("", fasilitasHandler.DeleteFasilitas)
	fasilitas.GET("single", fasilitasHandler.FetchFasilitasById)
	fasilitas.GET("foreignid", fasilitasHandler.FetchFasilitasByForeignID)

	penginapan := v1.Group("penginapan")
	penginapan.POST("", penginapanHandler.AddPenginapan)
	penginapan.GET("", penginapanHandler.FetchPenginapan)
	penginapan.PUT("", penginapanHandler.UpdatePenginapanByID)
	penginapan.DELETE("", penginapanHandler.DeletePenginapanByID)
	penginapan.GET("single", penginapanHandler.FetchPenginapanById)
	penginapan.GET("slug", penginapanHandler.FetchPenginapanBySlug)
	penginapan.GET("pagination", penginapanHandler.FetchPagination)
	penginapan.GET("search", penginapanHandler.FetchSearchPagination)

	user := v1.Group("user")
	user.POST("", userHandler.AddUser)
	user.GET("", userHandler.FetchAllUser)
	user.PUT("", userHandler.UpdateUserById)
	user.DELETE("", userHandler.DeleteUserById)
	user.GET("single", userHandler.FetchUserById)
	user.GET("username", userHandler.FetchUserByUsername)
	user.GET("email", userHandler.FetchUserByEmail)
	user.GET("pagination", userHandler.FetchPagination)
	user.GET("search", userHandler.FetchSearchPagination)

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
