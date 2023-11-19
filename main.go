// Importing necessary packages.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PlayerInfo struct is created to store user input.
type PlayerInfo struct {
	Name string
}

// QuizItem struct is created to store questions and answers.
type QuizItem struct {
	Question string
	Options  []string
	Answer   string
}

// The main function.
func main() {
	// Creating variables to store user input and configure the correctness, wrongness with integer type.
	var playerInput PlayerInfo
	var correctAnswers, incorrectAnswers int

	// Asking for player's name.
	fmt.Println() // Empty line for separation.
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	playerInput.Name = strings.TrimSpace(name)
	fmt.Println()

	// Collection of questions with a slice.
	questions := []QuizItem{
		{"What is the capital of France?", []string{"A. Paris", "B. Berlin", "C. Rome", "D. Madrid"}, "A"},
		{"Who is the author of 'To Kill a Mockingbird'?", []string{"A. J.K. Rowling", "B. Harper Lee", "C. George Orwell", "D. Ernest Hemingway"}, "B"},
		{"What is the square of 9?", []string{"A. 64", "B. 81", "C. 100", "D. 121"}, "B"},
	}

	// Iteration concept using range on a slice in Go.
	for _, quiz := range questions {
		fmt.Println(quiz.Question)
		for _, option := range quiz.Options {
			fmt.Println(option)
		}

		// Reading input from the player through the terminal or console.
		fmt.Print("Answer (A/B/C/D): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// Configuration of correctness and wrongness, if correct then +1 or vice versa.
		if strings.EqualFold(answer, quiz.Answer) {
			correctAnswers++
		} else {
			incorrectAnswers++
		}

		// Empty line for separation (<br>).
		fmt.Println()
	}

	// Displaying final score.
	fmt.Println("Quiz Statistics")
	fmt.Println("Name           : ", playerInput.Name)
	fmt.Println("Score          : ", correctAnswers)
	fmt.Println("Correct Answers: ", correctAnswers)
	fmt.Println("Incorrect Answers: ", incorrectAnswers)
}
