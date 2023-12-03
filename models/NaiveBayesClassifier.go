package models

import (
	"math"
	"sort"
	"strings"
)

type NaiveBayesClassifier struct {
	WordFreqs      map[string]map[string]map[string]float64 // kategoria -> gatunek -> słowo -> częstość
	CategoryProb   map[string]float64                       // kategoria -> prawdopodobieństwo
	VocabularySize float64
	CategoryCounts map[string]float64 // kategoria -> liczba słów
}
type genreScore struct {
	genre string
	score float64
}

func (nbc *NaiveBayesClassifier) Train(games []Game) {
	nbc.WordFreqs = make(map[string]map[string]map[string]float64)
	nbc.CategoryCounts = make(map[string]float64)

	for _, game := range games {
		words := strings.Fields(game.Description)
		updateWordFrequencies := func(genres []string, category string) {
			if nbc.WordFreqs[category] == nil {
				nbc.WordFreqs[category] = make(map[string]map[string]float64)
			}
			for _, genre := range genres {
				if nbc.WordFreqs[category][genre] == nil {
					nbc.WordFreqs[category][genre] = make(map[string]float64)
				}
				for _, word := range words {
					word = strings.ToLower(word)
					nbc.WordFreqs[category][genre][word]++
					nbc.CategoryCounts[category]++
				}
			}
		}

		// Aktualizuj częstotliwości dla każdej kategorii
		updateWordFrequencies(game.Genres.Basic, "Basic")
		updateWordFrequencies(game.Genres.Perspective, "Perspective")
		updateWordFrequencies(game.Genres.Topic, "Topic")
		updateWordFrequencies(game.Genres.Setting, "Setting")
	}
}

func (nbc *NaiveBayesClassifier) Predict(description string) Genres {
	words := strings.Fields(description)
	predictGenre := func(category string) []string {
		genreScores := make(map[string]float64)
		for genre := range nbc.WordFreqs[category] {
			prob := 0.0 // Zamiast używać math.Log(nbc.CategoryProb[category])
			for _, word := range words {
				word = strings.ToLower(word)
				// Wygładzanie Laplace'a
				wordProb := (nbc.WordFreqs[category][genre][word] + 1) / (nbc.CategoryCounts[category] + nbc.VocabularySize)
				prob += math.Log(wordProb) // Logarytm zwiększa numeryczną stabilność
			}
			genreScores[genre] = prob
		}
		topN := 3
		if category == "Perspective" {
			topN = 1
		}
		return getTopGenres(genreScores, topN, category)
	}

	return Genres{
		Basic:       predictGenre("Basic"),
		Perspective: predictGenre("Perspective"),
		Topic:       predictGenre("Topic"),
		Setting:     predictGenre("Setting"),
	}
}

func getTopGenres(genreScores map[string]float64, topN int, category string) []string {
	var scores []genreScore
	for genre, score := range genreScores {
		scores = append(scores, genreScore{genre, score})
	}

	// Sortowanie gatunków według wyników (od najwyższego do najniższego)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	// Wybieranie top N gatunków dla innych kategorii
	topGenres := make([]string, 0, topN)
	for i := 0; i < len(scores) && i < topN; i++ {
		topGenres = append(topGenres, scores[i].genre)
	}

	return topGenres
}
