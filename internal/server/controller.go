package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pklaudat/trello-sdk/internal/trello"
)

func GetBoardsController(ctx *gin.Context) {
	client := ctx.MustGet("trelloClient").(*trello.TrelloClient)

	name := ctx.Query("name")

	boards, err := client.GetBoardsHandler(name)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch boards: "})
		return
	}

	ctx.JSON(200, boards)
}

func CreateBoardController(ctx *gin.Context) {
	client := ctx.MustGet("trelloClient").(*trello.TrelloClient)

	var payload trello.BoardRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	client.CreateBoardHandler(&payload)

	ctx.JSON(201, gin.H{
		"message": "Board created successfully",
		"data":    payload,
	})
}

func UpdateBoardController(ctx *gin.Context) {

}

func DeleteBoardController(ctx *gin.Context) {}

func GetCardsController(ctx *gin.Context) {}

func CreateCardController(ctx *gin.Context) {}

func UpdateCardController(ctx *gin.Context) {}

func DeleteCardController(ctx *gin.Context) {}
