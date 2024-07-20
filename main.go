package main

import (
	models "cf-problem-finder/model"
	services "cf-problem-finder/service"
	"fmt"
	"sort"
)

// var TargetRatings = []int{800, 1000, 1200, 1400, 1600, 1900}

func main() {
	userHandles := services.ReadUserHandles()
	fmt.Printf("Succesfully parced %d users!\n", len(userHandles))

	TargetRatings := services.ReadTargetRatings()
	fmt.Printf("Succesfully parced %d target ratings!\n", len(TargetRatings))
	fmt.Printf("Finding problems rated at: %v\n", TargetRatings)

	// assert that len(TargetRatings) > 0
	if len(TargetRatings) == 1 && TargetRatings[0] == 0 {
		panic("No proper target ratings found!")
	}

	allProblems := services.FetchAllProblems()
	fmt.Printf("# of all Problems:\t\t%d\n", len(allProblems))

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
	fmt.Printf("Successfully picked %d/%d Problems!\n", len(pickedProblems), len(TargetRatings))
	for _, rating := range TargetRatings {
		problem, found := pickedProblems[rating]
		if !found {
			fmt.Printf("%d - No problem found!\n", rating)
		} else {
			fmt.Printf("%d - Problem %s : %s\n", rating, problem.Id(), problem.Name)
			fmt.Println(problem)
		}
		fmt.Println("---")
	}
}
