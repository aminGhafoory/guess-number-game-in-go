package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//create a random number
func randomInt(Max int) int {
	Source := rand.NewSource(time.Now().UnixNano())
	RandGen := rand.New(Source)
	return RandGen.Intn(Max)
}

//read User guess
func ReadUserGuess() int {
	var guess int
	fmt.Scanln(&guess)
	return guess
}

// simple Error handler
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	RandomNumber := randomInt(100)
	//fmt.Println(RandomNumber)
	success := false
	guessCounter := 0
	for !success {
		fmt.Print("Enter your guess:")
		userguess := ReadUserGuess()
		guessCounter += 1
		if userguess == RandomNumber {
			fmt.Printf("Success you guessed in %d guesses", guessCounter)

			success = true
		} else if userguess > RandomNumber {
			fmt.Println("your guess is bigger than Secret Number")
		} else if userguess < RandomNumber {
			fmt.Println("your guess is lower than Secret Number")
		}
	}
}
