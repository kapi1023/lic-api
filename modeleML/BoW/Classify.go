package bow

import (
	"strings"

	"github.com/kapi1023/licencjat/api/models"
	"github.com/kapi1023/licencjat/api/utils"
)

type GenreKeywords struct {
	Basic       map[string][]string
	Perspective map[string][]string
	Topic       map[string][]string
	Setting     map[string][]string
}

func Classify(trainingData []models.Game, newGameDescription string) models.Genres {
	wordDictionary := createDictionary(trainingData)
	newGameWords := strings.Fields(strings.ToLower(newGameDescription))

	// Tworzenie wektora BoW dla nowego opisu
	newGameVector := make([]int, len(wordDictionary))
	for i, word := range newGameWords {
		if wordDictionary[word] {
			newGameVector[i] = 1
		}
	}

	bowVectors := createVectors(trainingData, wordDictionary)

	categoryScores := make(map[string]int)
	topGenres := models.Genres{}

	for category, vectors := range bowVectors {
		for _, vector := range vectors {
			score := 0
			for i := 0; i < len(vector); i++ {
				score += vector[i] * newGameVector[i]
			}
			categoryScores[category] += score
		}
	}

	// Sortowanie kategorii według wyników (malejąco)
	sortedCategories := make([]string, len(categoryScores))
	i := 0
	for cat := range categoryScores {
		sortedCategories[i] = cat
		i++
	}

	for i := 0; i < len(sortedCategories); i++ {
		for j := i + 1; j < len(sortedCategories); j++ {
			if categoryScores[sortedCategories[i]] < categoryScores[sortedCategories[j]] {
				sortedCategories[i], sortedCategories[j] = sortedCategories[j], sortedCategories[i]
			}
		}
	}

	// Wybierz top 3 kategorie z "Basic", "Topic" i "Setting"
	for _, cat := range sortedCategories {
		if len(topGenres.Basic) < 3 && utils.Contains(topGenres.Basic, cat) == false {
			topGenres.Basic = append(topGenres.Basic, cat)
		}
		if len(topGenres.Perspective) < 3 && utils.Contains(topGenres.Perspective, cat) == false {
			topGenres.Perspective = append(topGenres.Perspective, cat)
		}
		if len(topGenres.Topic) < 3 && utils.Contains(topGenres.Topic, cat) == false {
			topGenres.Topic = append(topGenres.Topic, cat)
		}
		if len(topGenres.Setting) < 3 && utils.Contains(topGenres.Setting, cat) == false {
			topGenres.Setting = append(topGenres.Setting, cat)
		}
	}

	return topGenres

}
