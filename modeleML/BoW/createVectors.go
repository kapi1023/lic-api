package bow

import (
	"strings"

	"github.com/kapi1023/licencjat/api/models"
)

func createVectors(trainingData []models.Game, wordDictionary map[string]bool) map[string][][]int {
	bowVectors := make(map[string][][]int)

	for _, game := range trainingData {
		words := strings.Fields(strings.ToLower(game.Description))
		vector := make([]int, len(wordDictionary))

		for i, word := range words {
			if wordDictionary[word] {
				vector[i] = 1
			}
		}

		// Dodawanie wektora BoW tylko do odpowiednich kategorii
		for _, genre := range game.Genres.Basic {
			bowVectors[genre] = append(bowVectors[genre], vector)
		}
		for _, genre := range game.Genres.Perspective {
			bowVectors[genre] = append(bowVectors[genre], vector)
		}
		for _, genre := range game.Genres.Topic {
			bowVectors[genre] = append(bowVectors[genre], vector)
		}
		for _, genre := range game.Genres.Setting {
			bowVectors[genre] = append(bowVectors[genre], vector)
		}
	}

	return bowVectors
}
