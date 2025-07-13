package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BoardEntity struct {
	Id          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
}

type BoardRequest struct {
	Name        string `json: "name"`
	Description string `json: "desc"`
}

const (
	BoardsPath = "/boards"
)

func (t *TrelloClient) GetBoardsHandler(name string) ([]BoardEntity, error) {

	boardPath := fmt.Sprintf("%s?key=%s&token=%s", BoardsPath, t.APIKey, t.APIToken)

	if name != "" {
		boardPath = fmt.Sprintf("%s?name=%s&key=%s&token=%s", BoardsPath, name, t.APIKey, t.APIToken)

	}

	url := t.Endpoint + boardPath

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return []BoardEntity{}, err
	}

	resp, err := t.HTTPClient.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Printf("Error trying to fetch the boards %v", err)
		return []BoardEntity{}, err
	}

	defer resp.Body.Close()

	var result []BoardEntity

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return []BoardEntity{}, err
	}

	return result, nil
}

func (t *TrelloClient) CreateBoardHandler(payload *BoardRequest) {
	url := t.Endpoint + BoardsPath + t.APIKey + t.APIToken

	bodyBytes := bytes.Buffer{}
	if err := json.NewEncoder(&bodyBytes).Encode(payload); err != nil {
		return
	}
	req, err := http.NewRequest("POST", url, &bodyBytes)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := t.HTTPClient.Do(req)

	if err != nil || resp.StatusCode != http.StatusCreated {
		fmt.Println("Error creating board:", err)
		return
	}
}

func (t *TrelloClient) UpdateBoardHandler() {

}

func (t *TrelloClient) DeleteBoardHandler() {}
