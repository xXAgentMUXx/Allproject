package hangmanweb

import (
	hg "hangmanweb/Hangman-classic/functions"
	"html/template"
	"net/http"
	"strings"
)

type Game struct { // All variable for the struct of the game
	WordToGuess    string
	AttemptsLeft   int
	GuessedLetters []rune
	HiddenWord     []rune
	Message        string
	Finished       bool
	HardMode       bool
	VowelCount     int
}

var GameInstance Game

// function to define the function for the web
func Web(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/game.html"))             // Render the templates in the web
	hiddenWord := ConvertWordToString(GameInstance.HiddenWord)             // Render the hidden word
	guessedLetters := strings.ToUpper(string(GameInstance.GuessedLetters)) // Render the letter guessed by the user
	hangmanStage := hg.DrawHangman(10 - GameInstance.AttemptsLeft)         // Render the hangman in web

	data := struct { // All variable for the templates web
		Game
		HiddenWordStr     string
		GuessedLettersStr string
		HangmanStr        string
	}{
		Game:              GameInstance,
		HiddenWordStr:     hiddenWord,
		GuessedLettersStr: guessedLetters,
		HangmanStr:        hangmanStage,
	}
	tmpl.Execute(w, data) // Execute the response from user that do a request
}

// function to create the menu
func MenuHandler(w http.ResponseWriter, r *http.Request) { // Function to render the menu
	tmpl := template.Must(template.ParseFiles("web/home.html")) // Render the templates of menu in the web
	tmpl.Execute(w, nil)
}
func GameModeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("web/game_mode.html"))
    tmpl.Execute(w, nil)
}
// Function to handle quitting the game
func QuitHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/menu", http.StatusSeeOther) // Redirect to the main menu
}
// Function to display the credits
func Credit(w http.ResponseWriter, r *http.Request) { 
	tmpl := template.Must(template.ParseFiles("web/credits.html")) // Render the templates of credits in the web
	tmpl.Execute(w, nil)
}
// function to handle the normale rules
func RulesHandler_normal(w http.ResponseWriter, r *http.Request) { 
	tmpl := template.Must(template.ParseFiles("web/rules.html")) // Render the templates of menu in the web
	tmpl.Execute(w, nil)
}
// function to play the game
func GuessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		userInput := r.FormValue("input") //
		game := &GameInstance

		if len(userInput) == 1 { // if user enter one letter
			letter := []rune(userInput)[0]
			game.Message = hg.SubmitLetter(letter, []rune(game.WordToGuess), &game.HiddenWord, game.HardMode, &game.VowelCount, &game.AttemptsLeft, &game.GuessedLetters)
		} else {
			if strings.EqualFold(userInput, game.WordToGuess) { // if user enter two word or try to guess the word
				game.HiddenWord = []rune(game.WordToGuess)
				game.Finished = true
				game.Message = "Félicitations ! Vous avez deviné le mot."
			} else {
				game.AttemptsLeft -= 2
				if game.AttemptsLeft <= 0 {
					game.Finished = true
					game.Message = "Vous avez perdu ! Le mot était " + game.WordToGuess
				} else {
					game.Message = "Mauvais mot ! Vous perdez 2 tentatives."
				}
			}
		}

		// Check if the word is guessed letter by letter
		if string(game.HiddenWord) == game.WordToGuess {
			game.Finished = true
			game.Message = "Félicitations ! Vous avez deviné le mot."
		} else if game.AttemptsLeft <= 0 {
			game.Finished = true
			game.Message = "Vous avez perdu ! Le mot était " + game.WordToGuess
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
// Function to restart the game
func RestartGame(w http.ResponseWriter, r *http.Request) {
	hardMode := r.URL.Query().Get("mode") == "hard"            // URL to go to the hardmode
	word, _ := hg.SelectRandomWord("Hangman-classic/word.txt") // Select a random word if the game restart
	hiddenWord := hg.InitializeGame(word, hardMode)            // Initialize the hiddenword or the game

	GameInstance = Game{ // Add the value on variable when the game is restarted
		WordToGuess:    word,
		HiddenWord:     hiddenWord,
		AttemptsLeft:   10,
		GuessedLetters: []rune{},
		Finished:       false,
		HardMode:       hardMode,
		VowelCount:     0,
	}
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect the request from user to restart the game
}

// Function to convert word from rune to string
func ConvertWordToString(hiddenWord []rune) string {
	var displayWord string
	for _, r := range hiddenWord {
		if r == 0 {
			displayWord += "_"
		} else {
			displayWord += string(r) + " "
		}
	}
	return displayWord
}
