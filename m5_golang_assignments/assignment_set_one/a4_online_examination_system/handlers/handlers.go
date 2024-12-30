package handlers

import (
	"a4_online_examination_system/services"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RunQuiz starts the quiz and manages user interactions
func RunQuiz() {
	scanner := bufio.NewScanner(os.Stdin)
	questions := services.GetQuestions()

	fmt.Println("\nWelcome to the Quiz!")
	fmt.Println("Enter the number of your answer or type 'exit' to quit.")

	for i, question := range questions {
		services.DisplayQuestion(question, i)
		fmt.Print("\nYour answer: ")

		for {
			if !scanner.Scan() {
				return
			}

			input := strings.TrimSpace(scanner.Text())

			if strings.ToLower(input) == "exit" {
				fmt.Println("\nQuiz terminated early.")
				return
			}

			var answer int
			_, err := fmt.Sscanf(input, "%d", &answer)
			if err != nil || answer < 1 || answer > len(question.Options) {
				fmt.Print("Invalid input. Please enter a number between 1 and ",
					len(question.Options), "\nYour answer: ")
				continue
			}

			if services.EvaluateAnswer(question, answer) {
				fmt.Println("Correct!")
			} else {
				fmt.Printf("Incorrect. The correct answer was: %s\n",
					question.Options[question.CorrectAnswer-1])
			}
			break
		}
	}

	totalQuestions := len(questions)
	finalScore := services.GetFinalScore()
	performance := services.EvaluatePerformance(totalQuestions)

	fmt.Printf("\nQuiz completed!")
	fmt.Printf("\nFinal Score: %d/%d", finalScore, totalQuestions)
	fmt.Printf("\nPerformance Rating: %s\n", performance)
}
