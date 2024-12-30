package services

import (
	"a4_online_examination_system/models"
	"fmt"
)

var questions = []models.Question{
	{
		Text:          "What is the capital of France?",
		Options:       []string{"London", "Berlin", "Paris", "Madrid"},
		CorrectAnswer: 3,
	},
	{
		Text:          "Which programming language was created by Google?",
		Options:       []string{"Python", "Go", "Java", "Ruby"},
		CorrectAnswer: 2,
	},
	{
		Text:          "What is 2 + 2?",
		Options:       []string{"3", "4", "5", "6"},
		CorrectAnswer: 2,
	},
}

var score int

func DisplayQuestion(question models.Question, questionNum int) {
	fmt.Printf("\nQuestion %d: %s\n", questionNum+1, question.Text)
	for i, option := range question.Options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
}

func EvaluateAnswer(question models.Question, userAnswer int) bool {
	if userAnswer == question.CorrectAnswer {
		score++
		return true
	}
	return false
}

func GetFinalScore() int {
	return score
}

func EvaluatePerformance(totalQuestions int) string {
	percentage := float64(score) / float64(totalQuestions) * 100
	switch {
	case percentage >= 90:
		return "Excellent"
	case percentage >= 70:
		return "Good"
	case percentage >= 50:
		return "Fair"
	default:
		return "Needs Improvement"
	}
}

func GetQuestions() []models.Question {
	return questions
}
