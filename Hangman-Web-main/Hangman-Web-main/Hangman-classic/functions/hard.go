package functions

import (
	"strings"
	"math/rand"
	"unicode"
)

// Initialize the game
func InitializeGame(word string, hardMode bool) []rune {
	wordRunes := []rune(word) // Convert word in runes
   
	var revealedLetters int
	if hardMode {
	 revealedLetters = len(wordRunes)/3 - 1 // reveal letter in hard mode
	} else {
	 revealedLetters = len(wordRunes)/2 - 1 // reveal letter in normal mode
	}
   
	hiddenWordRunes := make([]rune, len(wordRunes)) // render the _ in the word
	for i := 0; i < len(wordRunes); i++ {
	 hiddenWordRunes[i] = '_' 
	}
   
	revealed_Indices := make(map[int]bool)  // variable for reveal letter in the beginning of the game
   
	for len(revealed_Indices) < revealedLetters { // select randomly the letter to reveal
	 randomIndex := rand.Intn(len(wordRunes)) 
   
	 if !revealed_Indices[randomIndex] { // if letter are not in the revealed then do not revealed
	  hiddenWordRunes[randomIndex] = wordRunes[randomIndex]
	  revealed_Indices[randomIndex] = true
   
		   }
	   }
	return hiddenWordRunes // return the hidden word with random letters revealed
   }

// Submit a letter
func SubmitLetter(guessRune rune, normalizedWord []rune, hiddenWord *[]rune, hardMode bool, vowelCount *int, AttemptsLeft *int, guessedLetters *[]rune) string {
	vowels := "AEOUIYaeouiy"
	correctGuess := false
	penaltyApplied := false
	message := ""

	// Normalize letter in upper word
	guessRune = unicode.ToUpper(guessRune)
	normalizedGuess := normalizeLetter(guessRune)

	// Check if letter is already guessed
	for _, guessed := range *guessedLetters {
		if guessed == normalizedGuess {
			if hardMode {
				*AttemptsLeft--
				message = "Lettre déjà utilisée. Vous perdez une tentative."
			} else {
				message = "Lettre déjà utilisée."
			}
			return message
		}
	}

	// add the used letter in the guess letter
	*guessedLetters = append(*guessedLetters, normalizedGuess)

	// take all the letter and it's variant
	for i, r := range normalizedWord {
		// Normalize letter in the upper word
		if normalizeLetter(unicode.ToUpper(r)) == normalizedGuess && (*hiddenWord)[i] == '_' {
			(*hiddenWord)[i] = r // Reveal the corretc word (whenever if it's a variant or not)
			correctGuess = true
		}
	}

	// Add the vowelcount if the letter is a vowel
	if strings.ContainsRune(vowels, normalizedGuess) {
		*vowelCount++
	}

	// Apply penality if the user used 3 vowel or more
	if correctGuess && strings.ContainsRune(vowels, normalizedGuess) && hardMode && *vowelCount >= 3 && !penaltyApplied {
		message = "Pénalité appliquée pour avoir utilisé plus de 3 voyelles."
		*AttemptsLeft--
		penaltyApplied = true
	}

	//if there is penality and the letter is incorrect, then it's 1 attempsleft
	if !correctGuess && !penaltyApplied {
		*AttemptsLeft--
		message = "Lettre incorrecte."
	}

	// Check if the word is totally guessed
	if string(*hiddenWord) == string(normalizedWord) {
		message = "Félicitations ! Vous avez deviné le mot."
		for i := range *hiddenWord {
			(*hiddenWord)[i] = normalizedWord[i] // Display all the letter in the hiddenword
		}
	}

	return message
}

// Function to add the variants letters
func normalizeLetter(r rune) rune {
	switch r {
	case 'É', 'È', 'Ê', 'Ë':
		return 'E'
	case 'Á', 'À', 'Â', 'Ä':
		return 'A'
	case 'Î', 'Ï':
		return 'I'
	case 'Ô', 'Ö':
		return 'O'
	case 'Ú', 'Ù', 'Û', 'Ü':
		return 'U'
	case 'Ç':
		return 'C'
	default:
		return r // return if there is no accent
	}
}