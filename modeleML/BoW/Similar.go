package bow

import (
	"strings"

	"github.com/kapi1023/licencjat/api/models"
)

func CreateDictionaryNowe(games []models.Game) map[string]int {
	dictionary := make(map[string]int)
	for _, game := range games {
		words := strings.Fields(game.Description)
		for _, word := range words {
			dictionary[word]++
		}
	}
	return dictionary
}

// Funkcja tworząca wektor na podstawie opisu gry
func createVectorNowe(description string, dictionary map[string]int) []int {
	words := strings.Fields(description)
	vector := make([]int, len(dictionary))
	i := 0
	for word := range dictionary {
		for _, w := range words {
			if w == word {
				vector[i]++
			}
		}
		i++
	}
	return vector
}

// Funkcja znajdująca najbardziej podobną grę
func FindMostSimilarGame(newGame string, games []models.Game, dictionary map[string]int) models.Genres {
	newGameVector := createVectorNowe(newGame, dictionary)
	maxSimilarity := 0
	var mostSimilarGame models.Genres
	for _, game := range games {
		gameVector := createVectorNowe(newGame, dictionary)
		similarity := 0
		for i := range newGameVector {
			similarity += newGameVector[i] * gameVector[i]
		}
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			mostSimilarGame = game.Genres
		}
	}
	return mostSimilarGame
}
