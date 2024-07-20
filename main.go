package main

import (
	models "cf-problem-finder/model"
	services "cf-problem-finder/service"
	"fmt"
	"sort"
)

var TargetRatings = []int{800, 1000, 1200, 1400, 1600, 1900}

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

	pickedProblems := make(map[int]models.Problem)
	for _, problem := range unsolvedProblems {
		for _, targetRating := range TargetRatings {
			if _, found := pickedProblems[targetRating]; found {
				continue
			}
			if problem.Rating == targetRating {
				pickedProblems[targetRating] = problem
				break
			}
		}
		if len(pickedProblems) == len(TargetRatings) {
			break
		}
	}
	fmt.Printf("Successfully picked %d Problems!\n", len(pickedProblems))
	for _, rating := range TargetRatings {
		problem := pickedProblems[rating]
		fmt.Printf("%d - Problem %s: %s\n", rating, problem.Id(), problem.Name)
		fmt.Println(problem)
		fmt.Println("---")
	}

}
