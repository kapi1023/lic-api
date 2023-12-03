package levenshtein

import (
	"sort"
	"strings"

	"github.com/kapi1023/licencjat/api/models"
	"github.com/kapi1023/licencjat/api/utils"
)

func Classify(trainingData []models.Game, newGameDescription string) models.Genres {
	threshold := 3 // Próg odległości Levenshteina do akceptacji

	categoryScores := make(map[string]int)
	topGenres := models.Genres{}

	for _, game := range trainingData {
		distance := distance(strings.ToLower(game.Description), strings.ToLower(newGameDescription))
		if distance <= threshold {
			// Dodaj do wyników kategorii
			for _, genre := range game.Genres.Basic {
				categoryScores[genre]++
			}
			for _, genre := range game.Genres.Perspective {
				categoryScores[genre]++
			}
			for _, genre := range game.Genres.Topic {
				categoryScores[genre]++
			}
			for _, genre := range game.Genres.Setting {
				categoryScores[genre]++
			}
		}
	}

	// Sortowanie kategorii według wyników (malejąco)
	sortedCategories := sortCategoryScores(categoryScores)

	for i := 0; i < len(sortedCategories); i++ {
		for j := i + 1; j < len(sortedCategories); j++ {
			if categoryScores[sortedCategories[i]] < categoryScores[sortedCategories[j]] {
				sortedCategories[i], sortedCategories[j] = sortedCategories[j], sortedCategories[i]
			}
		}
	}

	// Dodawanie gatunków do topGenres z zastosowaniem filtrów
	for _, cat := range sortedCategories {
		if len(topGenres.Basic) < 3 && !utils.Contains(topGenres.Basic, cat) {
			topGenres.Basic = append(topGenres.Basic, cat)
		}
		if len(topGenres.Topic) < 3 && !utils.Contains(topGenres.Topic, cat) {
			topGenres.Topic = append(topGenres.Topic, cat)
		}
		if len(topGenres.Setting) < 3 && !utils.Contains(topGenres.Setting, cat) {
			topGenres.Setting = append(topGenres.Setting, cat)
		}
		if len(topGenres.Perspective) < 1 && !utils.Contains(topGenres.Perspective, cat) {
			topGenres.Perspective = append(topGenres.Perspective, cat)
		}
	}

	return topGenres
}

func sortCategoryScores(categoryScores map[string]int) []string {
	sortedCategories := make([]string, 0, len(categoryScores))
	for cat := range categoryScores {
		sortedCategories = append(sortedCategories, cat)
	}
	sort.Slice(sortedCategories, func(i, j int) bool {
		return categoryScores[sortedCategories[i]] > categoryScores[sortedCategories[j]]
	})
	return sortedCategories
}
