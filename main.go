package main

import (
	models "cf-problem-finder/model"
	services "cf-problem-finder/service"
	"fmt"
	"sort"
)

func main() {
	allProblems := services.FetchAllProblems()
	fmt.Printf("# of all Problems:\t\t%d\n", len(allProblems))

	userHandles := []string{"NapSaq", "mrmoon"}

	solvedProblems := services.FetchSolvedProblemsForUsers(userHandles)

	fmt.Printf("# of solved Problems:\t\t%d\n", len(solvedProblems))

	unsolvedProblems := []models.Problem{}
	for _, problem := range allProblems {
		if !solvedProblems[problem.Id()] {
			unsolvedProblems = append(unsolvedProblems, problem)
		}
	}

	fmt.Printf("# of unsolved Problems:\t\t%d\n", len(unsolvedProblems))

	sort.Slice(unsolvedProblems, func(i, j int) bool {
		return unsolvedProblems[i].SolvedCount > unsolvedProblems[j].SolvedCount
	})

	fmt.Println(unsolvedProblems[2])

}
