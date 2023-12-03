package naivebayes

import "github.com/kapi1023/licencjat/api/models"

func NewNaiveBayesClassifier() *models.NaiveBayesClassifier {
	return &models.NaiveBayesClassifier{
		WordFreqs:      make(map[string]map[string]map[string]float64),
		CategoryProb:   make(map[string]float64),
		CategoryCounts: make(map[string]float64),
		VocabularySize: 0,
	}
}
