package models

type Game struct {
	GameId      int
	Description string
	Title       string
	Genres      Genres
}

type Genres struct {
	Basic       []string
	Perspective []string
	Topic       []string
	Setting     []string
}
