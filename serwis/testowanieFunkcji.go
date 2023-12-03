package serwis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	bow "github.com/kapi1023/licencjat/api/modeleML/BoW"
	levenshtein "github.com/kapi1023/licencjat/api/modeleML/Levenshtein"
	naivebayes "github.com/kapi1023/licencjat/api/modeleML/NaiveBayes"
	predictest "github.com/kapi1023/licencjat/api/modeleML/PredicTest"
	"github.com/kapi1023/licencjat/api/models"
	"github.com/kapi1023/licencjat/api/utils"
)

const (
	testowySring = `Battlefield 3 leaps ahead of the competition with the power of Frostbite™ 2, the next installment of DICE's cutting-edge game engine. This state-of-the-art technology is the foundation on which Battlefield 3 is built, delivering superior visual quality, a grand sense of scale, massive destruction, dynamic audio and incredibly lifelike character animations. As bullets whiz by, walls crumble, and explosions throw you to the ground, the battlefield feels more alive and interactive than ever before. In Battlefield 3, players step into the role of the elite U.S. Marines where they will experience heart-pounding single player missions and competitive multiplayer actions ranging across diverse locations from around the globe including Paris, Tehran and New York.`
)

func testowanieFunkcji() {
	data, err := ioutil.ReadFile("games.json")
	if err != nil {
		fmt.Println("Błąd przy odczytywaniu pliku JSON:", err)
		return
	}
	var games []models.Game
	err = json.Unmarshal(data, &games)
	if err != nil {
		fmt.Println("Błąd przy dekodowaniu JSON:", err)
		return
	}

	genreKeywordsData, err := ioutil.ReadFile("keywords.json")
	if err != nil {
		fmt.Println("Błąd przy odczytywaniu pliku JSON:", err)
		return
	}
	var genreKeywords models.GenreKeywords
	err = json.Unmarshal(genreKeywordsData, &genreKeywords)
	if err != nil {
		fmt.Println("Błąd przy dekodowaniu JSON:", err)
		return
	}
	wynik := models.Genres{}
	var processedGames []models.Game
	// Usuwanie tagów HTML
	// Usuwanie słów z blacklisty
	// Aktualizacja opisu gry
	processedGames = filtracjaWynikow(games, processedGames)
	nbc := naivebayes.NewNaiveBayesClassifier()
	nbc.Train(processedGames)
	fmt.Println("NaiveBayes")
	wynik = nbc.Predict(testowySring)
	fmt.Println(wynik)
	fmt.Println("Predic")
	wynik = predictest.PredictGenres(testowySring, genreKeywords)
	fmt.Println(wynik)
	fmt.Println("FindMostSimilarGame")
	wynik = bow.FindMostSimilarGame(testowySring, processedGames, bow.CreateDictionaryNowe(processedGames))
	fmt.Println(wynik)
	fmt.Println("BoW")
	wynik = bow.Classify(processedGames, testowySring)
	fmt.Println(wynik)
	fmt.Println("levenshtein")
	wynik = levenshtein.Classify(processedGames, testowySring)
	fmt.Println(wynik)
}

func filtracjaWynikow(games []models.Game, processedGames []models.Game) []models.Game {
	for _, game := range games {

		cleanDescription := utils.RemoveHTMLTags(game.Description)

		words := strings.Fields(cleanDescription)
		var filteredWords []string
		for _, word := range words {
			if _, exists := utils.CommonWordsBlacklist[word]; !exists {
				filteredWords = append(filteredWords, word)
			}
		}

		game.Description = strings.Join(filteredWords, " ")
		processedGames = append(processedGames, game)
	}
	return processedGames
}
