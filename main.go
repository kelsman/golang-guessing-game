package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){


	for matching := false; !matching; {
	// generate secret number;

	secret := getRandomNumber();


	// Get user input to guess
	guess := getUserInput();
  matching = isMatching(secret, guess)

	}
	
}


func getRandomNumber() int{

	rand.Seed(time.Now().Unix())
	return rand.Int() % 11;
}

func getUserInput() int {
		fmt.Println("Please input your guesse");
		var input int;
		 _ , err := fmt.Scan(&input);
		 if err != nil {
			fmt.Println("Failed to parse your input")
		 }else {
			fmt.Println("You guessed:", input)
		 }
		 return input
}

func isMatching(secret , guess int) bool{

	if guess < secret {
		fmt.Println("Too low!")
		return false
	}else if guess > secret{
		fmt.Println("Too large")
		return false
	}else {
		fmt.Println("You guessed right!")
		return true
	}


}