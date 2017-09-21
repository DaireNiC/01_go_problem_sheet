//Author: Daire Ní Chatháin
//Problem sheet 1 : Guessing Game
/* A guessing game where the user must guess a secret number.
After every guess the program tells the user whether their number
was too large or too small. At the end the number of tries
is printed. It counts only as one try if the user inputs
the same number multiple times consecutively.
 */

package main

import "fmt"
import "math/rand"

func main() {
	
	var answer, guess, previous_guess, try_count int = 0,0,0,0

	//generate a pseudorandom number between 1 - 100
	answer = rand.Intn(100)

	//Take in user's guess
	fmt.Println("Take a guess: ")
    fmt.Scanf("%d", &guess)

	//Keep looping until user guesses correct num
	for guess != answer {

		fmt.Println("Take another guess: ")
    	fmt.Scanf("%d", &guess)

		//print "too high" or "too low" depending on input
		if guess < answer{
			fmt.Println("Too low!")
		}else if guess > answer{
			fmt.Println("Too high!")
		}

		//Update number of tries if same num entered consecutively 
		if guess == previous_guess {
			try_count ++
		}
		//update previous guess
		previous_guess = guess
	}
	//This confirms user guessed correctly
	if guess == answer{
		fmt.Println("Congratulations, you guessed correctly!")
		fmt.Println("\n=== Game stats ==== \n Number of tries:", try_count)
	}
	
}