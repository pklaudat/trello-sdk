package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pklaudat/trello-sdk/docs"
	"github.com/pklaudat/trello-sdk/internal/trello"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Middleware(trelloClient *trello.TrelloClient) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Set("trelloClient", trelloClient)
		ctx.Next()
	}
}

func InitializeRoutes(key, token, host string) {
	basePath := "/api/v1"

	docs.SwaggerInfo.Title = "Trello GO API"
	docs.SwaggerInfo.Description = "Trello GO Rest API"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	trelloClient := trello.NewTrelloClient(key, token, host)

	router.Use(Middleware(trelloClient))

	v1 := router.Group(basePath)
	{

		v1.GET("/boards:id", GetBoardsController)
		v1.POST("/boards", CreateBoardController)
		v1.PUT("/boards/:id", UpdateBoardController)
		v1.DELETE("/boards/:id", DeleteBoardController)
		v1.GET("/cards", GetCardsController)
		v1.POST("/cards", CreateCardController)
		v1.PUT("/cards/:id", UpdateCardController)
		v1.DELETE("/cards/:id", DeleteCardController)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
