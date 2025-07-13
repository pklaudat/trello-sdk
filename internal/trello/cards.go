package trello

type Card struct {
	Name string
	Id   string
}

const (
	CardsPath = "/cards"
)

func (t *TrelloClient) GetCardsHandler() {}

func (t *TrelloClient) CreateCardHandler() {}

func (t *TrelloClient) UpdateCardHandler() {}

func (t *TrelloClient) DeleteCardHandler() {}
