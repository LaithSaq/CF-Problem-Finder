package main

import (
	models "cf-problem-finder/model"
	services "cf-problem-finder/service"
	"fmt"
)

func main() {
	allProblems := services.FetchAllProblems()
	fmt.Println(len(allProblems))

	mySolvedProblems := services.FetchSolvedProblemsForUser("NapSaq")
	fmt.Println(len(mySolvedProblems))

	solvedProblems := make(map[string]bool)
	for _, problem := range mySolvedProblems {
		solvedProblems[problem.Id()] = true
	}

	fmt.Println(len(solvedProblems))

	unsolvedProblems := []models.Problem{}
	for _, problem := range allProblems {
		if !solvedProblems[problem.Id()] {
			unsolvedProblems = append(unsolvedProblems, problem)
		}
	}

	fmt.Println(len(unsolvedProblems))
}
