package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pklaudat/trello-sdk/internal/trello"
)

// Get Board details
// @Summary Get board details
// @Description Fetch board details using the id
// @Tags boards
// @Accept json
// @Produce json
// @Param id query string false "boards search by id"
// @Router /api/v1/boards [get]
func GetBoardsController(ctx *gin.Context) {
	client := ctx.MustGet("trelloClient").(*trello.TrelloClient)

	id := ctx.Query("id")

	boards, err := client.GetBoardHandler(id)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to fetch boards: "})
		return
	}

	ctx.JSON(200, boards)
}

// Create Board godoc
// @Summary Create trello board
// @Description Create new trello board
// @Tags boards
// @Accept json
// @Produce json
// @Router /api/v1/boards [post]
func CreateBoardController(ctx *gin.Context) {
	client := ctx.MustGet("trelloClient").(*trello.TrelloClient)

	var payload trello.BoardRequest
	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Printf("Failed to parse request body: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newBoard, err := client.CreateBoardHandler(&payload)
	if err != nil {
		fmt.Printf("Failed to create board: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Board created successfully",
		"data":    newBoard,
	})
}

func UpdateBoardController(ctx *gin.Context) {

}

func DeleteBoardController(ctx *gin.Context) {}

func GetCardsController(ctx *gin.Context) {}

func CreateCardController(ctx *gin.Context) {}

func UpdateCardController(ctx *gin.Context) {}

func DeleteCardController(ctx *gin.Context) {}
