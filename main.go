package main

import (
	models "cf-problem-finder/model"
	services "cf-problem-finder/service"
	"fmt"
	"sort"
	"slices"
)

// var targetRatings = []int{800, 1000, 1200, 1400, 1600, 1900}

func main() {
	userHandles := services.ReadUserHandles()
	fmt.Printf("Succesfully parced %d users!\n", len(userHandles))

	targetRatings := services.ReadTargetRatings()
	fmt.Printf("Succesfully parced %d target ratings!\n", len(targetRatings))
	fmt.Printf("Finding problems rated at: %v\n", targetRatings)

	// assert that len(targetRatings) > 0
	if len(targetRatings) == 1 && targetRatings[0] == 0 {
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

	pickedProblems := make([]models.Problem, len(targetRatings))
	pickedProblemsCount := 0

	for ratingIndex, targetRating := range targetRatings {
		for _, problem := range unsolvedProblems {
			if i := slices.IndexFunc(pickedProblems, func(p models.Problem) bool { return p.Id() == problem.Id() }); i != -1 {
				// skip if problem is already picked
				continue
			}
			if problem.Rating == targetRating {
				pickedProblems[ratingIndex] = problem
				pickedProblemsCount++
				break
			}
		}
	}
	fmt.Printf("Successfully picked %d/%d Problems!\n", pickedProblemsCount, len(targetRatings))
	for ratingIndex, rating := range targetRatings {
		problem := pickedProblems[ratingIndex]
		if problem.Id() == (models.Problem{}).Id() {
			fmt.Printf("%d - No problem found!\n", rating)
		} else {
			fmt.Printf("%d - Problem %s : %s\n", rating, problem.Id(), problem.Name)
			fmt.Println(problem)
		}
		fmt.Println("---")
	}
}
