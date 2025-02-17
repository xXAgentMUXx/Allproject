package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Define a flag for hard mode
	hardMode := flag.Bool("hard", false, "Activate hard mode (uses hardword.txt)")
	flag.Parse()

	var file *os.File
	var err error

	// Open the file based on the chosen mode
	if *hardMode {
		file, err = os.Open("hardword.txt") // Open hardword.txt if -h is tapped
	} else {
		file, err = os.Open("word.txt") // Else, open word.txt
	}

	if err != nil { // Handle the file opening error
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Close the file

	var randomWords []string // Create an array to store the words
	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line
	for scanner.Scan() { // Read the file until there are no more lines
		word := scanner.Text() // Store each word
		randomWords = append(randomWords, word) // Add each word to the array
	}

	// Handle errors during reading
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	if len(randomWords) == 0 {
		fmt.Println("No words available in the file")
		return
	}

	// Choose a random word
	rand.Seed(time.Now().UnixNano())
	word := randomWords[rand.Intn(len(randomWords))]

	lifes := 9

	// Create an array to store the hidden letters
	hidenWords := []string{}
	for range word {
		hidenWords = append(hidenWords, "_")
	}

	// Colors using ANSI escape codes
	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"

	// Print colored ASCII art of HANGMAN
	fmt.Println(green + `  _    _                                             |` + reset)
	fmt.Println(green + ` | |  | |                                            |` + reset)
	fmt.Println(green + ` | |__| | __ _ _ __   ___  _ ___ ___  ___  _ __      |   __  __  _ __    ______  __     __ ` + reset)
	fmt.Println(yellow + ` |  __  |/ _\` + "`" + `| '_ \ / _ \` + `| '_  '_` + `  |` + `/` + "`" + ` _\| '_ \     |   \ \/ / | '_ \  |  __  | \ \   / /` + reset)
	fmt.Println(yellow + ` | |  | | (_| | | | | (_| | | | | | | (_| | | | |    |    \  /  | | | | | |__| |  \ \_/ /` + reset)
	fmt.Println(red + ` |_|  |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|    |    /_/   |_| |_| |______|   \___/ ` + reset)
	fmt.Println(red + `                      __/ |                          |   ` + reset)
	fmt.Println(red + `                     |___/                           |   ` + reset)

	// Game loop
	for {
		// Show remaining lives and hidden word
		fmt.Printf("❤️ %s%d%s, Word: %s, Letter: ", red, lifes, reset, strings.Join(hidenWords, " "))
		var input string
		fmt.Scanln(&input)

		// Check if the player wants to quit
		if strings.ToLower(input) == "exit" {
			exitGame(word)
			return
		}

		input = strings.ToLower(input)

		foundWord := false // Set to false for every input round

		for i, letterWord := range word {
			if input == string(letterWord) { // Compare string with the letter of the word
				hidenWords[i] = string(letterWord)
				foundWord = true
			}
		}

		if !foundWord {
			lifes--
			printHangman(lifes) // Appelle la fonction pour afficher l'état du hangman après une erreur
		}

		// Game over conditions
		if lifes <= 0 {
			fmt.Printf("❤️ %s0%s, You lost, the word to find was: %s! \n", red, reset, word)
			break
		}

		if word == strings.Join(hidenWords, "") {
			fmt.Printf("❤️ %s%d%s, Congratulations, you won! The word was: %s! \n", green, lifes, reset, word)
			break
		}
	}
}

// exitGame displays an exit message
func exitGame(word string) {
	fmt.Printf("You exited the game. The word to find was: %s.\n", word)
}

// printHangman displays the hangman based on the number of lives left
func printHangman(lifes int) {
	fmt.Println("|---------|")
	fmt.Println("|         |")

	switch lifes {
	case 9:
		fmt.Println("|         0")
	case 8:
		fmt.Println("|         0")
		fmt.Println("|        /")
	case 7:
		fmt.Println("|         0")
		fmt.Println("|        /|")
	case 6:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
	case 5:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
		fmt.Println("|        /")
	case 4:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
		fmt.Println("|        / \\")
	case 3:
		fmt.Println("|        0")
		fmt.Println("|       /|\\")
		fmt.Println("|       / \\")
	case 2:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
		fmt.Println("|        / \\")
	case 1:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
		fmt.Println("|        / \\")
	case 0:
		fmt.Println("|         0")
		fmt.Println("|        /|\\")
		fmt.Println("|        / \\")
	}
	fmt.Println("|")
}
