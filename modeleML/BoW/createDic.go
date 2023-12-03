package bow

import (
	"strings"

	"github.com/kapi1023/licencjat/api/models"
)

func createDictionary(trainingData []models.Game) map[string]bool {
	wordDictionary := make(map[string]bool)

	for _, game := range trainingData {
		words := strings.Fields(strings.ToLower(game.Description))
		for _, word := range words {
			wordDictionary[word] = true
		}
	}

	return wordDictionary
}
