package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BoardResponse struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Desc           string  `json:"desc"`
	ShortURL       string  `json:"shortUrl"`
	URL            string  `json:"url"`
	IDOrganization string  `json:"idOrganization"`
	IDEnterprise   *string `json:"idEnterprise"` // use *string because it can be null
}

type BoardRequest struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

const (
	BoardsPath = "/boards"
)

func (t *TrelloClient) GetBoardHandler(id string) (BoardResponse, error) {
	boardPath := fmt.Sprintf("/1/boards/%s?key=%s&token=%s", id, t.APIKey, t.APIToken)
	url := t.Endpoint + boardPath

	fmt.Printf("GET board details for id: %s\n", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return BoardResponse{}, err
	}

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return BoardResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return BoardResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response BoardResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return BoardResponse{}, err
	}

	return response, nil
}

func (t *TrelloClient) CreateBoardHandler(payload *BoardRequest) (BoardResponse, error) {
	url := t.Endpoint + "/1" + BoardsPath + "?key=" + t.APIKey + "&token=" + t.APIToken

	fmt.Printf("New request to create a board %s\n", payload.Name)

	bodyBytes := bytes.Buffer{}
	if err := json.NewEncoder(&bodyBytes).Encode(payload); err != nil {
		return BoardResponse{}, err
	}

	req, err := http.NewRequest("POST", url, &bodyBytes)

	if err != nil {
		fmt.Println("Error creating request:", err)
		return BoardResponse{}, err
	}

	resp, err := t.HTTPClient.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Error creating board:", err)
		return BoardResponse{}, err
	}

	defer resp.Body.Close()

	fmt.Printf("Response content: %v", resp.Body)

	var createdBoard BoardResponse

	if err := json.NewDecoder(resp.Body).Decode(&createdBoard); err != nil {
		fmt.Println("Error decoding response:", err)
		return BoardResponse{}, err
	}

	return createdBoard, nil
}

func (t *TrelloClient) UpdateBoardHandler() {

}

func (t *TrelloClient) DeleteBoardHandler() {}
