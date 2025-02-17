package main

import (
	"net/http"
	"hangmanweb"
	hg "hangmanweb/Hangman-classic/functions"
)

func main() {
	word, _ := hg.SelectRandomWord("Hangman-classic/word.txt") // Select a random word in the web

hiddenWord := make([]rune, len(word)) // Render the hidden word 
for i := range hiddenWord {
    hiddenWord[i] = '_'  
}

hangmanweb.GameInstance = hangmanweb.Game{ // Add in the html the variable
    WordToGuess:    word,
    HiddenWord:     hiddenWord,
    AttemptsLeft:   10,
    GuessedLetters: []rune{},
    Finished:       false,
	HardMode:       false,
	VowelCount:      0,
}
	http.HandleFunc("/menu", hangmanweb.MenuHandler)       //http://localhost:8080/menu (to go in the menu)
	http.HandleFunc("/quit", hangmanweb.QuitHandler)
	http.HandleFunc("/rules", hangmanweb.RulesHandler_normal)     //http://localhost:8080/rules (to go in the rules of the game)
	http.HandleFunc("/restart", hangmanweb.RestartGame)	   //http://localhost:8080/restart (to restart a new game, to restart in hardmode : http://localhost:8080//restart?mode=hard)
	http.HandleFunc("/", hangmanweb.Web) 				   //http://localhost:8080/        (to go to the game without passing the menu)
	http.HandleFunc("/game", hangmanweb.GameModeHandler) //http://localhost:8080/game      (to go to the game mode)
	http.HandleFunc("/guess", hangmanweb.GuessHandler)
	http.HandleFunc("/credits", hangmanweb.Credit)         // http://localhost:8080/credits
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web")))) // take the file css and execute them with the html
	http.ListenAndServe(":8080", nil) // Add a port for the request
}
