package main

import (
	"fmt"
	"math/rand"
	//"regexp"
)

var bodyPart = []string{"head", "left arm", "right arm", "left leg", "right leg", "body"}
var words = []string{"bananas", "apples", "peach", "elephant", "orange", "burritos", "ithaca college"}

var currentBody []string
var guesses []string //create an empty set to hold all guessed letters that user entered
var word string

func drawBodyPart(wrongGuessCount int) []string {
	for i := 0; i <= wrongGuessCount; i++ {
		currentBody = append(currentBody, bodyPart[i])
	}
	return currentBody
}

func addGuessToSet(guess string) string {
	//when letters is guessed then automatically add to set
	//if it already exists then return "enter another letter"
	var duplicate bool = false
	for _, i := range guesses {
		if i == guess {
			duplicate = true
		}
	}
	if duplicate {
		return "You guessed a duplicate letter, guess another.\n"
	} else {
		guesses = append(guesses, guess)
		return "You guessed " + guess + "\n"
	}
}

func chooseRandomWord() string {
	var randomWord = rand.Intn(len(words))

	var pick string = words[randomWord]
	return pick
}

func getWordProgress() (string, bool) { //TODO
	//init wordProgress with a string of underscores of length that matches the word
	//look through guesses - if a guessed letter is in word, then look for where it exists in the word and
	//replace the respective underscores in wordProgress. Might want to make this a set?
	//return wordProgress and if the guessed letter was in the word or not
	//regexp package may be useful?

	return "", false
}

func main() {

	var wrongGuessCount int
	var answer string

	word = chooseRandomWord() //randomly select word user has to guess
	fmt.Println("Your word is :  ")
	getWordProgress()

	for answer != "quit" {
		//TODO Ask if they want to guess a letter or the whole word, then prompt for it
		fmt.Scanf("%s", &answer)

		if answer == "letter" {
			fmt.Println("\nGuess a letter or enter 'quit' to stop: ")
			fmt.Scanf("%s", &answer) //TODO need to do some user error checking here, make sure they only enter a letter

			addGuessToSet(answer)
			wordProgress, guessedRight := getWordProgress()

			if !guessedRight {
				wrongGuessCount++
				currentBody = drawBodyPart(wrongGuessCount)
			}

			fmt.Println("Word Progress: " + wordProgress)
			fmt.Printf("Bodyparts visible: %v", currentBody) //this prints the whole list?
		}
	}
}