package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pklaudat/trello-sdk/internal/trello"
)

func Middleware(trelloClient *trello.TrelloClient) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		ctx.Set("trelloClient", trelloClient)
		ctx.Next()
	}
}

func InitializeRoutes(key, token, host string) {
	basePath := "/api/v1"
	router := gin.Default()

	trelloClient := trello.NewTrelloClient(key, token, host)

	router.Use(Middleware(trelloClient))

	v1 := router.Group(basePath)
	{
		v1.GET("/boards", GetBoardsController)
		v1.POST("/boards", CreateBoardController)
		v1.PUT("/boards/:id", UpdateBoardController)
		v1.DELETE("/boards/:id", DeleteBoardController)
		v1.GET("/cards", GetCardsController)
		v1.POST("/cards", CreateCardController)
		v1.PUT("/cards/:id", UpdateCardController)
		v1.DELETE("/cards/:id", DeleteCardController)

	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
