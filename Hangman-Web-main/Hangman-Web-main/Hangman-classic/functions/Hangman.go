package functions

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode/utf8"
)

// Select a word randomly
func SelectRandomWord(mot string) (string, error) {
	file, err := os.Open(mot) // open file
	if err != nil {
		return "Fichier non trouvé ou manquant", err
	}
	defer file.Close() // check if file is close in the end of function

	words := []string{}
	scanner := bufio.NewScanner(file) // add a scanner
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	random := rand.Intn(len(words)) // choose a random word
	return words[random], nil
}

// Play the game
func PlayHangman(hardMode bool) {
	menu() // call menu

	for {
		fmt.Print("\n\n\n\n\n\n\n\n")

		word, err := SelectRandomWord("Hangman-classic/word.txt") // call selectRandomWord
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier:", err) // end if there is an error
			return
		}

		hiddenWord := []rune(InitializeGame(word, hardMode))
		normalizedWord := []rune(word)     // convert the word in runes for the accents
		vowelCount := 0                    // number of vowel
		AttemptsLeft := 10
		stageOfDeath := 0                  // case of hangman
		lifes := 10                        // number of life
		usedLetters := make(map[rune]bool) // letter that has been used
		guessedLetters := make([]rune, 0)
		var message string                 // display the message to the player

		for stageOfDeath < lifes && string(hiddenWord) != string(normalizedWord) { // render of the game
			DrawHangman(stageOfDeath)
			fmt.Println("Tapez 'exit' si vous voulez quitter")
			fmt.Println("Mot à deviner:", string(hiddenWord))
			if !hardMode {
				fmt.Println("Lettres devinées:", string(guessedLetters))
			}
			fmt.Print("Choisissez une lettre ou devinez le mot (en mettant les accents si besoin): ")
			fmt.Printf("Il vous reste %d tentatives.\n", lifes-stageOfDeath)

			if message != "" { // display message for each condition
				fmt.Println(message)
				message = ""
			}

			var guess string
			fmt.Scanln(&guess) // read the imput of user

			guess = strings.ToUpper(guess) // convert letter in  capital letter

			if strings.ToLower(guess) == "exit" { // leave the game
				fmt.Println("Vous avez quitté le jeu.")
				return
			}

			if len(guess) > 1 {
				if guess == word { // if the word is correct
					hiddenWord = []rune(word)
				} else { //else remove 2 lifes
					message = "Mauvais Mot ! Vous perdez 2 vies."
					stageOfDeath += 2
					AttemptsLeft -= 2 // reduce the attempts (web)
				}
				continue
			}

			guessRune, _ := utf8.DecodeRuneInString(guess)    // decode the hidden word in rune
			normalizedGuessRune := normalizeLetter(guessRune) // Normalize rune for the accents

			if usedLetters[normalizedGuessRune] {
				// En mode normal, juste un message, sans perte de tentative
				message = "Lettre déjà utilisée."
				if hardMode {
					message = "Lettre déjà utilisée. Vous perdez une tentative."
					stageOfDeath++
					AttemptsLeft-- // Reduce the attempts (web)
				}
				continue
			}
			
			usedLetters[normalizedGuessRune] = true

			// if mode normal then  it display the letter guessed
			if !hardMode {
				guessedLetters = append(guessedLetters, []rune(guess)...)
			}

			// call the function submit letter
			message = SubmitLetter(guessRune, normalizedWord, &hiddenWord, hardMode, &vowelCount, &AttemptsLeft, &guessedLetters)

		}
		// if you win or lose, then it display this message
		if string(hiddenWord) == string(normalizedWord) {
			fmt.Println("Félicitations ! Vous avez deviné le mot :", string(normalizedWord))
		} else {
			DrawHangman(lifes) // final case of hangman
			fmt.Println("Vous avez perdu ! Le mot était :", string(normalizedWord))
		}

		var replay string
		fmt.Print("Appuyez sur 'Entrée' pour rejouer ou tapez 'exit' pour quitter: ")
		fmt.Scanln(&replay)                    // replay the game
		if strings.ToLower(replay) == "exit" { // leave the game
			fmt.Println("Merci d'avoir joué !")
			break
		}
	}
}

// function to display the hangman
func DrawHangman(stage_of_death int) string {

	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n") // SPace for each guess in the game

	if stage_of_death == 0 { // For not to display the first case of hangman
		return ""
	}

	file, err := os.Open("Hangman-classic/hangman.txt") // Open the file hangman.txt
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return "Erreur lors de la lecture du fichier du pendu."
	}
	defer file.Close()

	hangman := []string{}
	scanner := bufio.NewScanner(file)

	// Scan the file
	for scanner.Scan() {
		hangman = append(hangman, scanner.Text())
	}

	// Each case of 8 lines
	stageSize := 8

	if stage_of_death == 0 {
		return ""
	}

	start := (stage_of_death - 1) * stageSize 
	end := start + stageSize

	if start >= len(hangman) { // index for the lenght of hangman
		return ""
	}
	if end > len(hangman) {
		end = len(hangman)
	}
	return strings.Join(hangman[start:end],"\n")
}

// function to display the menu
func menu() {
	// add colors in the menu
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	// render the menu
	fmt.Println(green + `  _    _                                             |` + reset)
	fmt.Println(green + ` | |  | |                                            |` + reset)
	fmt.Println(green + ` | |__| | __ _ _ __   ___  _ ___ ___  ___  _ __      |   __  __  _ __    ______  __     __ ` + reset)
	fmt.Println(yellow + ` |  __  |/ _\` + "`" + `| '_ \ / _ \` + `| '_  '_` + `  |` + `/` + "`" + ` _\| '_ \     |   \ \/ / | '_ \  |  __  | \ \   / /` + reset)
	fmt.Println(yellow + ` | |  | | (_| | | | | (_| | | | | | | (_| | | | |    |    \  /  | | | | | |__| |  \ \_/ /` + reset)
	fmt.Println(red + ` |_|  |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|    |    /_/   |_| |_| |______|   \___/ ` + reset)
	fmt.Println(red + `                      __/ |                          |   ` + reset)
	fmt.Println(red + `                     |___/                           |   ` + reset)

	fmt.Println("Appuyez sur 'Entrée' pour commencer le jeu ou tapez 'exit' pour quitter")

	reader := bufio.NewReader(os.Stdin) // begin the game
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "exit" { // leave in the end of the game
		fmt.Println("Vous avez quitté le jeu.")
		os.Exit(0)
	}
}