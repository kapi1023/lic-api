package predictest

import (
	"sort"
	"strings"

	"github.com/kapi1023/licencjat/api/models"
)

type genreScore struct {
	genre string
	score int
}

func PredictGenres(description string, genreKeywords models.GenreKeywords) models.Genres {
	var predictedGenres models.Genres
	wordCounts := make(map[string]int)

	// Rozdziel opis na słowa i policz ich wystąpienia
	words := strings.Fields(description)
	for _, word := range words {
		wordCounts[strings.ToLower(word)]++
	}

	// Funkcja pomocnicza do przewidywania gatunku na podstawie słów kluczowych
	predictGenre := func(keywords map[string]map[string]int) []string {
		var genreScores []genreScore
		for genre, keywordsMap := range keywords {
			score := 0
			for keyword, weight := range keywordsMap {
				score += wordCounts[keyword] * weight
			}
			if score > 0 {
				genreScores = append(genreScores, genreScore{genre, score})
			}
		}

		// Sortuj gatunki według wyniku
		sort.Slice(genreScores, func(i, j int) bool {
			return genreScores[i].score > genreScores[j].score
		})

		// Wybierz trzy gatunki z najwyższymi wynikami
		topGenres := []string{}
		for i := 0; i < len(genreScores) && i < 3; i++ {
			topGenres = append(topGenres, genreScores[i].genre)
		}
		return topGenres
	}

	// Przewidywanie dla każdej kategorii
	predictedGenres.Basic = predictGenre(genreKeywords.Basic)
	predictedGenres.Perspective = predictGenre(genreKeywords.Perspective)
	predictedGenres.Topic = predictGenre(genreKeywords.Topic)
	predictedGenres.Setting = predictGenre(genreKeywords.Setting)

	return predictedGenres
}
