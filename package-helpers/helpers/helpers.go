package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	randno "math/rand"
	"strconv"
)

func RandomStringGenerator(len int) (string, error) {
	// get some random bytes

	randomBytes := make([]byte, len)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	// encode the random bytes
	// to string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString[:len], nil
}

// returns one
// or
// multiple random words from
// list

func ReturnRandomWords(options ...int) (string, []string, error) {

	var word string
	var words []string
	var err error
	var limit int = 1

	if len(options) > 0 {
		limit = options[0]
	}

	dictionary := []string{"Aboideau", "Abord", "Adhibit", "Agastopia", "Ambiguous", "Aurora",
		"Barmcloths", "Bastinade", "Bungalow", "Bypass", "Cachinnate", "Caffoy", "Calvary", "Capricious",
		"Chads", "Cherish", "Cynical", "Dap", "Deliquescent", "Decalcomania", "Demure", "Dragoman", "Eclaircise",
		"Eftsoons", "Elixir", "Esoteric", "Euphoria", "Flabbergast", "Flimflam", "Floccinaucinihilipilification",
		"Gallizes", "Guise", "Halfpace", "Hent", "Hospitium", "Hypnosis", "Idyllic", "Impignorate", "Immure",
		"Incendiary", "Jejune", "Jentacular", "Kakorrhaphiophobia", "Levant", "Limerence", "Lithest", "Loquacious",
		"Lucid", "Mabble", "Melancholy", "Mellifluous", "Misanthrope", "Narrative", "Nebulize", "Nemesis", "Nostalgic",
		"Nudiustertian", "Obdurate", "Omnishambles", "Opulence", "Paradox", "Pejorate", "Persiflage", "Perspicacious",
		"Pique", "Plethora", "Pristine", "Quackles", "Quell", "Quincunx", "Quintessence", "Quire", "Quixotic", "Rambunctious",
		"Renaissance", "Sanguine", "Scrumptious", "Serene", "Serendipity", "Snood", "Tenacious", "Tintinnabulation",
		"Tittynopes", "Triskaidekaphobia", "Ulotrichous", "Uncanny", "Uranism", "Vamp", "Velleity", "Wanderlust", "Wanton",
		"Wherewithal", "Winklepicker", "Xenophobia", "Xertzes", "Yarborough", "Yawner", "Yearn", "Zealot", "Zenith"}

	if limit == 1 {
		// grab a random key from the dictionary
		// keep in mind, we are generating a numeric
		// value based on the limit of the dictionary
		// in other words, it's length
		word = dictionary[randno.Intn(len(dictionary))]
	} else {

		// since this is a slice, we can grab data from random intervals
		// pick a random point for iterating

		// shuffle the slice first
		for i := len(dictionary) - 1; i > 0; i-- {
			j := randno.Intn(i + 1)
			dictionary[i], dictionary[j] = dictionary[j], dictionary[i]
		}

		if len(dictionary) > limit {
			rand_point := len(dictionary) - limit

			for i := rand_point; i < len(dictionary); i++ {
				words = append(words, dictionary[i])
			}

		} else {
			err = errors.New("Word count cannot be larger than the dictionary. Current dictionary size is:" + strconv.Itoa(len(dictionary)))
		}

	}

	return word, words, err

}

// Counts String Occurence
// In Input Slice

func CountOccurenceInSlice(inputSlice []string) (map[string]int, error) {
	result := make(map[string]int)

	if len(inputSlice) > 0 {
		for _, string := range inputSlice {
			result[string]++
		}

		return result, nil
	}

	return nil, errors.New("we got a problem, the input contains 0 elements")
}

// keeps the window open untill the user presses ENTER
func Exit() {
	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
