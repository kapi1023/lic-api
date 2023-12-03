package bowv2

import (
	"sort"
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

func CategorizeGame(description string, keywords GenreKeywords) models.Genres {
	words := strings.Fields(strings.ToLower(description))
	score := make(map[string]int)

	// Liczenie wystąpień kluczowych słów dla każdej kategorii
	countKeywords(words, keywords.Basic, score)
	countKeywords(words, keywords.Perspective, score)
	countKeywords(words, keywords.Topic, score)
	countKeywords(words, keywords.Setting, score)

	// Wybieranie top kategorii
	return models.Genres{
		Basic:       topGenres(score, keywords.Basic, 3),
		Perspective: topGenres(score, keywords.Perspective, 1),
		Topic:       topGenres(score, keywords.Topic, 3),
		Setting:     topGenres(score, keywords.Setting, 3),
	}
}

func countKeywords(words []string, genreMap map[string][]string, score map[string]int) {
	for _, word := range words {
		for genre, keys := range genreMap {
			if utils.Contains(keys, word) {
				score[genre]++
			}
		}
	}
}

func topGenres(score map[string]int, genreMap map[string][]string, n int) []string {
	type genreScore struct {
		genre string
		score int
	}

	var genreScores []genreScore
	for genre := range genreMap {
		genreScores = append(genreScores, genreScore{genre, score[genre]})
	}

	// Sortowanie gatunków wg wyników
	sort.Slice(genreScores, func(i, j int) bool {
		return genreScores[i].score > genreScores[j].score
	})

	// Wybieranie top N gatunków
	topGenres := make([]string, 0, n)
	for i := 0; i < n && i < len(genreScores); i++ {
		topGenres = append(topGenres, genreScores[i].genre)
	}

	return topGenres
}
